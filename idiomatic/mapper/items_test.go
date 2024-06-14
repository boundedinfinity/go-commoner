package mapper_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/stretchr/testify/assert"
)

func Test_Items(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected []mapper.Item[string, int]
	}{
		{
			name:     "case 1",
			input:    map[string]int{"A": 1, "B": 2},
			expected: []mapper.Item[string, int]{{"A", 1}, {"B", 2}},
		},
		{
			name:     "case 2",
			input:    map[string]int{},
			expected: []mapper.Item[string, int]{},
		},
		{
			name:     "case 3",
			input:    nil,
			expected: []mapper.Item[string, int]{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := mapper.Items(tc.input)
			assert.ElementsMatch(tt, tc.expected, actual)
		})
	}
}
