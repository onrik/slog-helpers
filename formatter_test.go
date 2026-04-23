package sloghelpers

import "testing"

type testStringer struct {
}

func (testStringer) String() string {
	return "test"
}

type testStruct struct {
	ID   int64
	Name string
}

func TestFormatValue(t *testing.T) {
	t.Run("TestInt", func(t *testing.T) {
		value := 78
		f := FormatValue(value)
		if f != "78" {
			t.Fail()
		}
	})

	t.Run("TestPointer", func(t *testing.T) {
		value := 78
		f := FormatValue(&value)
		if f != "78" {
			t.Fail()
		}
	})

	t.Run("TestStringer", func(t *testing.T) {
		f := FormatValue(testStringer{})
		if f != "test" {
			t.Fail()
		}
	})

	t.Run("TestStruct", func(t *testing.T) {
		f := FormatValue(testStruct{
			ID:   43589,
			Name: "foo",
		})
		if f != "{ID:43589 Name:foo}" {
			t.Fail()
		}
	})
}
