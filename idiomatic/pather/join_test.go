package pather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	tcs := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "join 1",
			input:    []string{"a"},
			expected: "a",
		},
		{
			name:     "join 2",
			input:    []string{"a", "b"},
			expected: "a/b",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := pather.Join(tc.input...)
			assert.Equal(tt, actual, tc.expected)
		})
	}
}
