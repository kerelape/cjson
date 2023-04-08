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
	"encoding/json"
	"testing"

	cjson "github.com/kerelape/cjson/pkg"
)

func TestNumber_MarshalJSON(t *testing.T) {
	set := []struct {
		name    string
		subject cjson.Number
		want    []byte
	}{
		{
			name:    "zero",
			subject: cjson.NewNumber(0),
			want:    []byte("0"),
		},
		{
			name:    "simple integer",
			subject: cjson.NewNumber(123),
			want:    []byte("123"),
		},
		{
			name:    "simple fractional number",
			subject: cjson.NewNumber(0.2),
			want:    []byte("0.2"),
		},
		{
			name:    "simple fractional number with more than one digit",
			subject: cjson.NewNumber(4.235),
			want:    []byte("4.235"),
		},
		{
			name:    "fractional number with 'e'",
			subject: cjson.NewNumber(5.3e4),
			want:    []byte("53000"),
		},
		{
			name:    "fractional number with negative 'e'",
			subject: cjson.NewNumber(5.3e-4),
			want:    []byte("0.00053"),
		},
	}
	for _, c := range set {
		t.Run(c.name, func(t *testing.T) {
			actual, err := json.Marshal(c.subject)
			if err != nil {
				t.Errorf("Failed to marshal: %v", err)
			}
			if !bytes.Equal(actual, c.want) {
				t.Errorf("Expected %s, got %s", string(c.want), string(actual))
			}
		})
	}
}
