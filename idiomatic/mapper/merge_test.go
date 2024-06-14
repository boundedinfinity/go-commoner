package mapper_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/stretchr/testify/assert"
)

func Test_MergeCopy(t *testing.T) {
	testCases := []struct {
		name     string
		input1   map[string]int
		input2   map[string]int
		expected map[string]int
	}{
		{
			name:     "case 1",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
		},
		{
			name:     "case 2",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   map[string]int{},
			expected: map[string]int{"A": 1, "B": 2},
		},
		{
			name:     "case 3",
			input1:   map[string]int{},
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int{"C": 3, "D": 4},
		},
		{
			name:     "case 4",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   nil,
			expected: map[string]int{"A": 1, "B": 2},
		},
		{
			name:     "case 5",
			input1:   nil,
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int{"C": 3, "D": 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			// input1 := mapper.Copy(tc.input1)
			// input2 := mapper.Copy(tc.input2)
			actual := mapper.MergeCopy(tc.input1, tc.input2)

			assert.Equal(tt, tc.expected, actual, "expected == actual")
			// assert.Equal(tt, input1, tc.input1, "input1 == tc.input1")
			// assert.Equal(tt, input2, tc.input2, "input2 == tc.input2")
		})
	}
}

func Test_MergeInto(t *testing.T) {
	testCases := []struct {
		name     string
		input1   map[string]int
		input2   map[string]int
		expected map[string]int
	}{
		{
			name:     "case 1",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
		},
		{
			name:     "case 2",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   map[string]int{},
			expected: map[string]int{"A": 1, "B": 2},
		},
		{
			name:     "case 3",
			input1:   map[string]int{},
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int{"C": 3, "D": 4},
		},
		{
			name:     "case 4",
			input1:   map[string]int{"A": 1, "B": 2},
			input2:   nil,
			expected: map[string]int{"A": 1, "B": 2},
		},
		{
			name:     "case 5",
			input1:   nil,
			input2:   map[string]int{"C": 3, "D": 4},
			expected: map[string]int(nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			mapper.MergeInto(tc.input1, tc.input2)

			assert.Equal(tt, tc.expected, tc.input1)
		})
	}
}
