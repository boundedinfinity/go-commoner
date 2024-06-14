package mapper_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/stretchr/testify/assert"
)

func Test_Reject(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected map[string]int
		fn       func(k string, v int) bool
	}{
		{
			name:     "case 1",
			input:    map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			expected: map[string]int{"B": 2, "C": 3, "D": 4},
			fn:       func(k string, v int) bool { return k == "A" },
		},
		{
			name:     "case 2",
			input:    map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			expected: map[string]int{"A": 1, "C": 3, "D": 4},
			fn:       func(k string, v int) bool { return v == 2 },
		},
		{
			name:     "case 3",
			input:    map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			expected: map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			fn:       nil,
		},
		{
			name:     "case 4",
			input:    nil,
			expected: map[string]int{},
			fn:       func(k string, v int) bool { return v == 2 },
		},
		{
			name:     "case 5",
			input:    map[string]int{},
			expected: map[string]int{},
			fn:       func(k string, v int) bool { return v == 2 },
		},
		{
			name:     "case 6",
			input:    nil,
			expected: nil,
			fn:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := mapper.Reject(tc.input, tc.fn)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
