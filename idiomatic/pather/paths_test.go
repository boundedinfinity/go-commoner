package pather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
		err      error
		fn       func(string) string
	}{
		{
			name:     "base",
			input:    "a/b/c",
			expected: "c",
			fn:       pather.Paths.Base,
		},
		{
			name:     "dir",
			input:    "a/b/c",
			expected: "a/b",
			fn:       pather.Paths.Dir,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input)
			assert.Equal(t, actual, tc.expected)
		})
	}
}
