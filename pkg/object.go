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
	"fmt"
	"strconv"
	"strings"

	option "github.com/kerelape/option/pkg"
)

// Object is a JSON object.
type Object struct {
	content []objectEntry
}

// NewObject returns a new Object.
func NewObject() Object {
	return Object{[]objectEntry{}}
}

func (o Object) MarshalJSON() ([]byte, error) {
	entries := make([]string, 0, len(o.content))
	for _, entry := range o.content {
		valueBytes, err := entry.value.MarshalJSON()
		if err != nil {
			return nil, err
		}
		entries = append(entries, fmt.Sprintf("%s:%s", strconv.Quote(entry.key), string(valueBytes)))
	}
	return []byte(fmt.Sprintf("{%s}", strings.Join(entries, ","))), nil
}

func (o Object) String() option.Option[StringLeaf] {
	return option.None[StringLeaf]()
}

func (o Object) Number() option.Option[NumberLeaf] {
	return option.None[NumberLeaf]()
}

func (o Object) Boolean() option.Option[BooleanLeaf] {
	return option.None[BooleanLeaf]()
}

func (o Object) Object() option.Option[ObjectBranch] {
	return option.Some[ObjectBranch](o)
}

func (o Object) Array() option.Option[ArrayBranch] {
	return option.None[ArrayBranch]()
}

func (o Object) With(key string, value Node) ObjectBranch {
	content := make([]objectEntry, len(o.content))
	copy(content, o.content)
	for i, entry := range o.content {
		if entry.key == key {
			content[i].value = value
			return Object{content}
		}
	}
	return Object{append(content, objectEntry{key, value})}
}

func (o Object) Found(key string) option.Option[Node] {
	for _, entry := range o.content {
		if entry.key == key {
			return option.Some(entry.value)
		}
	}
	return option.None[Node]()
}

func (o Object) Keys() []string {
	keys := make([]string, 0, len(o.content))
	for _, entry := range o.content {
		keys = append(keys, entry.key)
	}
	return keys
}

type objectEntry struct {
	key   string
	value Node
}
