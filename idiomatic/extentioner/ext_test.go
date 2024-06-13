package extentioner_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/stretchr/testify/assert"
)

func Test_Ext(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
		fn       func(string) string
	}{
		{
			name:     "ext",
			input:    "/some/path/file.json",
			expected: ".json",
			fn:       extentioner.Ext,
		},
		{
			name:     "extOnly",
			input:    "/some/path/file.json",
			expected: "json",
			fn:       extentioner.ExtOnly,
		},
		{
			name:     "normalize 1",
			input:    "txt",
			expected: ".txt",
			fn:       extentioner.Normalize,
		},
		{
			name:     "normalize 1",
			input:    ".txt",
			expected: ".txt",
			fn:       extentioner.Normalize,
		},
		{
			name:     "strip",
			input:    "/some/path/file.json",
			expected: "/some/path/file",
			fn:       extentioner.Strip,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input)
			assert.Equal(tt, actual, tc.expected)
		})
	}
}

func Test_Swap_1(t *testing.T) {
	input := "/some/path/file.json"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, "json", "txt")

	assert.Equal(t, expected, actual)
}

func Test_Swap_2(t *testing.T) {
	input := "/some/path/file.json"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}

func Test_Swap_3(t *testing.T) {
	input := "/some/path/file.txt"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}
