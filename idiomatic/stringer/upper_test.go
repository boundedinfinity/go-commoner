package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_ToUpper_normal_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "string",
			expected: "STRING",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToUpper(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToUpper_typed_string(t *testing.T) {
	type MyString string

	testCases := []struct {
		name     string
		input    MyString
		expected string
	}{
		{
			name:     "case 1",
			input:    MyString("string"),
			expected: "STRING",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToUpper(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToUpperFirst_normal_string(t *testing.T) {
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToUpperFirst(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToUpperFirst_typed_string(t *testing.T) {
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToUpperFirst(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
