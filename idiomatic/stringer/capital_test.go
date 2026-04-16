package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_ToCapital_normal_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "string",
			expected: "String",
		},
		{
			name:     "case 2",
			input:    "string string",
			expected: "String String",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToCapital(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToCapital_typed_string(t *testing.T) {
	type MyString string

	testCases := []struct {
		name     string
		input    MyString
		expected string
	}{
		{
			name:     "case 1",
			input:    MyString("string"),
			expected: "String",
		},
		{
			name:     "case 2",
			input:    MyString("string string"),
			expected: "String String",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToCapital(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
