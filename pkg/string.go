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

// String is JSON string.
type String struct {
	content string
}

// NewString returns a new String.
func NewString(content string) String {
	return String{content}
}

func (s String) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(s.content)), nil
}

func (s String) Content() string {
	return s.content
}

func (s String) String() option.Option[StringLeaf] {
	return option.Some[StringLeaf](s)
}

func (s String) Number() option.Option[NumberLeaf] {
	return option.None[NumberLeaf]()
}

func (s String) Boolean() option.Option[BooleanLeaf] {
	return option.None[BooleanLeaf]()
}

func (s String) Object() option.Option[ObjectBranch] {
	return option.None[ObjectBranch]()
}

func (s String) Array() option.Option[ArrayBranch] {
	return option.None[ArrayBranch]()
}
