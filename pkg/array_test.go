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
	"github.com/stretchr/testify/assert"
)

func TestArray_MarshalJSON(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		subject := cjson.NewArray()
		want := []byte("[]")
		actual, err := json.Marshal(subject)
		if err != nil {
			t.Errorf("Failed to marshal: %v", err)
		}
		if !bytes.Equal(actual, want) {
			t.Errorf("Expected %s, got %s", string(want), string(actual))
		}
	})
	t.Run("one value", func(t *testing.T) {
		subject := cjson.NewArray(cjson.NewString("jojo"))
		want := []byte(`["jojo"]`)
		actual, err := json.Marshal(subject)
		if err != nil {
			t.Errorf("Failed to marshal: %v", err)
		}
		if !bytes.Equal(actual, want) {
			t.Errorf("Expected %s, got %s", string(want), string(actual))
		}
	})
	t.Run("multiple values", func(t *testing.T) {
		subject := cjson.NewArray(cjson.NewString("jojo"), cjson.NewNull())
		want := []byte(`["jojo",null]`)
		actual, err := json.Marshal(subject)
		if err != nil {
			t.Errorf("Failed to marshal: %v", err)
		}
		if !bytes.Equal(actual, want) {
			t.Errorf("Expected %s, got %s", string(want), string(actual))
		}
	})
}

func TestArray_Without(t *testing.T) {
	t.Run("no indexes", func(t *testing.T) {
		assert.Equal(
			t,
			3,
			cjson.NewArray(cjson.NewNull(), cjson.NewNull(), cjson.NewNull()).Without().Size(),
			"Without should not change the array if no values passed",
		)
	})
	t.Run("removes middle", func(t *testing.T) {
		assert.Equal(
			t,
			"Hello, World!",
			cjson.NewArray(cjson.NewNull(), cjson.NewNull(), cjson.NewString("Hello, World!")).
				Without(1).
				At(1).
				Value().
				String().
				Value().
				Content(),
			"Expected Without to remove the middle object",
		)
	})
}
