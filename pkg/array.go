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
	"errors"
	"fmt"
	"strings"

	option "github.com/kerelape/option/pkg"
)

var ErrOutOfBounds = errors.New("out of bounds")

// Array is a JSON array.
type Array struct {
	content []Node
}

// NewArray returns a new Array with content from source.
func NewArray(source ...Node) Array {
	return Array{source}
}

func (a Array) MarshalJSON() ([]byte, error) {
	entries := make([]string, 0, len(a.content))
	for _, value := range a.content {
		if value == nil {
			value = Null{}
		}
		valueBytes, err := value.MarshalJSON()
		if err != nil {
			return nil, err
		}
		entries = append(entries, string(valueBytes))
	}
	return []byte(fmt.Sprintf("[%s]", strings.Join(entries, ","))), nil
}

func (a Array) Content() []Node {
	return a.content
}

func (a Array) String() option.Option[StringLeaf] {
	return option.NewNoneReason[StringLeaf](ErrWrongType)
}

func (a Array) Number() option.Option[NumberLeaf] {
	return option.NewNoneReason[NumberLeaf](ErrWrongType)
}

func (a Array) Boolean() option.Option[BooleanLeaf] {
	return option.NewNoneReason[BooleanLeaf](ErrWrongType)
}

func (a Array) Object() option.Option[ObjectBranch] {
	return option.NewNoneReason[ObjectBranch](ErrWrongType)
}

func (a Array) Array() option.Option[ArrayBranch] {
	return option.NewSome[ArrayBranch](a)
}

func (a Array) With(values ...Node) ArrayBranch {
	return Array{append(a.content, values...)}
}

func (a Array) Size() int {
	return len(a.content)
}

func (a Array) At(index int) option.Option[Node] {
	if index > (a.Size()-1) || index < 0 {
		return option.NewNoneReason[Node](ErrOutOfBounds)
	}
	return option.NewSome(a.content[index])
}

func (a Array) Without(indexes ...int) ArrayBranch {
	if len(indexes) == 0 {
		return Array{a.content}
	}
	content := make([]Node, 0, len(a.content))
	index := indexes[len(indexes)-1]
	content = append(content, a.content[:index]...)
	content = append(content, a.content[index+1:]...)
	result := Array{content}
	return result.Without(indexes[:len(indexes)-1]...)
}
