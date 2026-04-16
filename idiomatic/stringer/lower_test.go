package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_ToLower_normal_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "STRING",
			expected: "string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToLower(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToLower_typed_string(t *testing.T) {
	type MyString string

	testCases := []struct {
		name     string
		input    MyString
		expected string
	}{
		{
			name:     "case 1",
			input:    MyString("STRING"),
			expected: "string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToLower(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToLowerFirst_normal_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "STRING",
			expected: "sTRING",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToLowerFirst(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ToLowerFirst_typed_string(t *testing.T) {
	type MyString string

	testCases := []struct {
		name     string
		input    MyString
		expected string
	}{
		{
			name:     "case 1",
			input:    MyString("STRING"),
			expected: "sTRING",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.ToLowerFirst(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
