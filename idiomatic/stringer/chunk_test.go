package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		size     int
		expected []string
	}{
		{
			name:     "Chunk 123456789 / 3",
			input:    "123456789",
			size:     3,
			expected: []string{"123", "456", "789"},
		},
		{
			name:     "Chunk 123456789 / 4",
			input:    "123456789",
			size:     4,
			expected: []string{"1234", "5678", "9"},
		},
		{
			name:     "Chunk 12 / 3",
			input:    "12",
			size:     3,
			expected: []string{"12"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.Chunk(tc.size, tc.input)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}
