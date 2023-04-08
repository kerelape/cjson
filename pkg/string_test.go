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

func TestString_MarshalJSON(t *testing.T) {
	t.Run("surrounds with quotes", func(t *testing.T) {
		subject := cjson.NewString("string")
		want := []byte(`"string"`)
		actual, err := subject.MarshalJSON()
		if err != nil {
			t.Errorf("Failed to marshal: %v", err)
		}
		if !bytes.Equal(actual, want) {
			t.Errorf(`Expected %s, got %s`, string(want), string(actual))
		}
	})
	t.Run("escapes quotation marks", func(t *testing.T) {
		subject := cjson.NewString(`They said "konnichiwa"`)
		want := []byte(`"They said \"konnichiwa\""`)
		actual, err := subject.MarshalJSON()
		if err != nil {
			t.Errorf("Failed to marshal: %v", err)
		}
		if !bytes.Equal(actual, want) {
			t.Errorf("Expected %s, got %s", string(want), string(actual))
		}
	})
}

func TestString_Content(t *testing.T) {
	t.Run("returns expected content", func(t *testing.T) {
		want := "abcdef"
		subject := cjson.NewString(want)
		actual := subject.Content()
		if actual != want {
			t.Errorf("Expected %s, got %s", want, actual)
		}
	})
}

func TestString_String(t *testing.T) {
	t.Run("returns correct object", func(t *testing.T) {
		want := "abcdef"
		subject := cjson.NewString(want)
		actual := subject.String()
		if actual.Empty() {
			t.Error("Expected to return non-empty value")
		}
		if actual.Value().Content() != want {
			t.Errorf("Expected %s, got %s", want, actual.Value().Content())
		}
	})
}

func TestString_Number(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if cjson.NewString("").Number().NotEmpty() {
			t.Error("Expected String to return an empty NumberValue")
		}
	})
}

func TestString_Boolean(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if cjson.NewString("").Boolean().NotEmpty() {
			t.Error("Expected String to return an empty BooleanValue")
		}
	})
}

func TestString_Object(t *testing.T) {
	t.Run("returns an empty object", func(t *testing.T) {
		if cjson.NewString("").Object().NotEmpty() {
			t.Error("Expected String to return an empty ObjectValue")
		}
	})
}

func TestString_Array(t *testing.T) {
	t.Run("returns an empty array", func(t *testing.T) {
		if cjson.NewString("").Array().NotEmpty() {
			t.Error("Expected String to return an empty ArrayValue")
		}
	})
}
