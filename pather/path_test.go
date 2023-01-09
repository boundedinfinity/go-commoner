package pather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/stretchr/testify/assert"
)

type testString string

func Test_Path(t *testing.T) {
	tcs := []struct {
		name     string
		input    testString
		expected string
		err      error
		fn       func(testString) string
	}{
		{
			name:     "base",
			input:    testString("a/b/c"),
			expected: "c",
			fn:       pather.Base[testString],
		},
		{
			name:     "dir",
			input:    testString("a/b/c"),
			expected: "a/b",
			fn:       pather.Dir[testString],
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input)
			assert.Equal(t, actual, tc.expected)
		})
	}
}
