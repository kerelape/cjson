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

import (
	"encoding/json"

	option "github.com/kerelape/option/pkg"
)

// Value is a JSON value.
type Value interface {
	json.Marshaler
}

// Node is JSON node.
type Node interface {
	Value

	// String returns this Node as String if possible.
	String() option.Option[StringLeaf]

	// Number returns this Node as Number if possible.
	Number() option.Option[NumberLeaf]

	// Boolean returns this Node as Boolean if possible.
	Boolean() option.Option[BooleanLeaf]

	// Object returns this Node as Object if possible.
	Object() option.Option[ObjectBranch]

	// Array returns this Node as Array if possible.
	Array() option.Option[ArrayBranch]
}

// Leaf is a final value in a JSON tree.
type Leaf[T any] interface {
	Value
	Content() T
}

// StringLeaf is a leaf that you can get a string from.
type StringLeaf Leaf[string]

// NumberLeaf is a leaf that you can get a float64 from.
type NumberLeaf Leaf[float64]

// BooleanLeaf is a leaf that you can get a bool from.
type BooleanLeaf Leaf[bool]

type ObjectBranch interface {
	Value
	With(key string, value Node) ObjectBranch
	Found(key string) option.Option[Node]
}

type ArrayBranch interface {
	Value
	With(values ...Node) ArrayBranch
	Size() int
	At(index int) option.Option[Node]
}
