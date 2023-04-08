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

func TestNull_MarshalJSON(t *testing.T) {
	subject := cjson.Null{}
	want := []byte("null")
	actual, err := subject.MarshalJSON()
	if err != nil {
		t.Errorf("Failed to marshal: %v", err)
	}
	if !bytes.Equal(actual, want) {
		t.Errorf("Expected %s, got %s", string(want), string(actual))
	}
}

func TestNull_String(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if (cjson.Null{}).String().NotEmpty() {
			t.Error("Expected Null to return an empty NumberValue")
		}
	})
}

func TestNull_Number(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if (cjson.Null{}).Number().NotEmpty() {
			t.Error("Expected Null to return an empty NumberValue")
		}
	})
}

func TestNull_Boolean(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if (cjson.Null{}).Boolean().NotEmpty() {
			t.Error("Expected Null to return an empty BooleanValue")
		}
	})
}

func TestNull_Object(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if (cjson.Null{}).Object().NotEmpty() {
			t.Error("Expected Null to return an empty ObjectValue")
		}
	})
}

func TestNull_Array(t *testing.T) {
	t.Run("returns an empty array", func(t *testing.T) {
		if (cjson.Null{}).Array().NotEmpty() {
			t.Error("Expected Null to return an empty ArrayValue")
		}
	})
}
