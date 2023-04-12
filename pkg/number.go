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

// Number is a JSON number.
type Number struct {
	content float64
}

// NewNumber returns a new Number.
func NewNumber(content float64) Number {
	return Number{content}
}

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(n.content, 'f', -1, 64)), nil
}

func (n Number) Content() float64 {
	return n.content
}

func (n Number) String() option.Option[StringLeaf] {
	return option.NewNoneReason[StringLeaf](ErrWrongType)
}

func (n Number) Number() option.Option[NumberLeaf] {
	return option.NewSome[NumberLeaf](n)
}

func (n Number) Boolean() option.Option[BooleanLeaf] {
	return option.NewNoneReason[BooleanLeaf](ErrWrongType)
}

func (n Number) Object() option.Option[ObjectBranch] {
	return option.NewNoneReason[ObjectBranch](ErrWrongType)
}

func (n Number) Array() option.Option[ArrayBranch] {
	return option.NewNoneReason[ArrayBranch](ErrWrongType)
}
