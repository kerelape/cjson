package cjson

import (
	"bytes"
	"encoding/json"
	"io"
)

// Document is a bridge between the json package and csjon.
type Document struct {
	root Node
}

func NewDocument() *Document {
	return &Document{nil}
}

func (d *Document) UnmarshalJSON(b []byte) error {
	root, err := Parse(b)
	if err != nil {
		return err
	}
	d.root = root
	return nil
}

// Root returns the root node of the Doucment.
func (d *Document) Root() Node {
	if d.root == nil {
		return NewNull()
	}
	return d.root
}

// Read reads JSON from source Reader.
func Read(source io.Reader) (Node, error) {
	root := any(nil)
	if err := json.NewDecoder(source).Decode(&root); err != nil {
		return nil, err
	}
	return Build(root), nil
}

// Parse parses JSON from source bytes.
func Parse(source []byte) (Node, error) {
	return Read(bytes.NewBuffer(source))
}

// Build builds a JSON tree.
func Build(node any) Node {
	if n, isNumber := node.(float64); isNumber {
		return NewNumber(n)
	}
	if s, isString := node.(string); isString {
		return NewString(s)
	}
	if b, isBoolean := node.(bool); isBoolean {
		return NewBoolean(b)
	}
	if o, isObject := node.(map[string]any); isObject {
		object := NewObject()
		for key, value := range o {
			object = object.With(key, Build(value)).(Object)
		}
		return object
	}
	if a, isArray := node.([]any); isArray {
		nodes := make([]Node, 0, len(a))
		for _, n := range a {
			nodes = append(nodes, Build(n))
		}
		return NewArray(nodes...)
	}
	return NewNull()
}
