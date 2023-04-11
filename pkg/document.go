// MIT License
//
// # Copyright (c) 2023 kerelape
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cjson

import "encoding/json"

// Document is a json document.
type Document struct {
	content []byte
}

// NewDocument creates a new Document.
func NewDocument(content []byte) Document {
	return Document{content}
}

// Parse parses the document.
func (d Document) Parse() (Node, error) {
	var content any
	if err := json.Unmarshal(d.content, &content); err != nil {
		return nil, err
	}
	return Build(content), nil
}

// Build builds a JSON tree.
func Build(node any) Node {
	if node == nil {
		return NewNull()
	}
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
	return NewInvalid()
}
