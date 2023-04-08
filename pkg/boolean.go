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
	"strconv"

	option "github.com/kerelape/option/pkg"
)

// Boolean is a JSON boolean.
type Boolean struct {
	content bool
}

// NewBoolean returns a new Boolean.
func NewBoolean(content bool) Boolean {
	return Boolean{content}
}

func (b Boolean) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(b.content)), nil
}

func (b Boolean) Content() bool {
	return b.content
}

func (b Boolean) String() option.Option[StringLeaf] {
	return option.None[StringLeaf]()
}

func (b Boolean) Number() option.Option[NumberLeaf] {
	return option.None[NumberLeaf]()
}

func (b Boolean) Boolean() option.Option[BooleanLeaf] {
	return option.Some[BooleanLeaf](b)
}

func (b Boolean) Object() option.Option[ObjectBranch] {
	return option.None[ObjectBranch]()
}

func (b Boolean) Array() option.Option[ArrayBranch] {
	return option.None[ArrayBranch]()
}
