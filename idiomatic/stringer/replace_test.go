package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Replace(t *testing.T) {
	testCases := []struct {
		name     string
		olds     []string
		new      string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			olds:     []string{"x"},
			new:      "_",
			input:    "x 1 x 2 x 3",
			expected: "_ 1 _ 2 _ 3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.Replace(tc.input, tc.new, tc.olds...)
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}
