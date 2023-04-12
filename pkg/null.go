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

import option "github.com/kerelape/option/pkg"

// Null is JSON null value.
type Null [0]any

// NewNull returns a new Null node.
func NewNull() Null {
	return Null{}
}

// MarshalJSON serializes Null to JSON null.
func (n Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (n Null) String() option.Option[StringLeaf] {
	return option.NewNone[StringLeaf]()
}

func (n Null) Number() option.Option[NumberLeaf] {
	return option.NewNone[NumberLeaf]()
}

func (n Null) Boolean() option.Option[BooleanLeaf] {
	return option.NewNone[BooleanLeaf]()
}

func (n Null) Object() option.Option[ObjectBranch] {
	return option.NewNone[ObjectBranch]()
}

func (n Null) Array() option.Option[ArrayBranch] {
	return option.NewNone[ArrayBranch]()
}
