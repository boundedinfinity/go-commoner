package mapper_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/stretchr/testify/assert"
)

func Test_Keys(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "case 1",
			input:    map[string]int{"A": 1, "B": 2},
			expected: []string{"A", "B"},
		},
		{
			name:     "case 2",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "case 3",
			input:    nil,
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := mapper.Keys(tc.input)
			assert.ElementsMatch(tt, tc.expected, actual)
		})
	}
}
