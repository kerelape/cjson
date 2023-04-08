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

package cjson_test

import (
	"bytes"
	"testing"

	cjson "github.com/kerelape/cjson/pkg"
)

func TestObject_MarshalJSON(t *testing.T) {
	cases := []struct {
		name    string
		subject cjson.Object
		want    []byte
	}{
		{
			name:    "empty",
			subject: cjson.NewObject(),
			want:    []byte("{}"),
		},
		{
			name:    "one value",
			subject: cjson.NewObject().With("a", cjson.Null{}).(cjson.Object),
			want:    []byte(`{"a":null}`),
		},
		{
			name:    "multiple values",
			subject: cjson.NewObject().With("a", cjson.Null{}).With("b", cjson.Null{}).With("c", cjson.Null{}).(cjson.Object),
			want:    []byte(`{"a":null,"b":null,"c":null}`),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.subject.MarshalJSON()
			if err != nil {
				t.Errorf("Failed to marshal: %v", err)
			}
			if !bytes.Equal(actual, c.want) {
				t.Errorf("Expected %s, got %s", string(c.want), string(actual))
			}
		})
	}
}
