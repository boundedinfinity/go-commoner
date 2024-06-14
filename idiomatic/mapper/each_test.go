package mapper_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/stretchr/testify/assert"
)

func Test_Each(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected int
		fn       func(initial *int) func(k string, v int)
	}{
		{
			name:     "case 1",
			input:    map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			expected: 10,
			fn:       func(initial *int) func(string, int) { return func(k string, v int) { *initial += v } },
		},
		{
			name:     "case 2",
			input:    map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
			expected: 0,
			fn:       nil,
		},
		{
			name:     "case 3",
			input:    nil,
			expected: 0,
			fn:       func(initial *int) func(string, int) { return func(k string, v int) { *initial += v } },
		},
		{
			name:     "case 4",
			input:    map[string]int{},
			expected: 0,
			fn:       func(initial *int) func(string, int) { return func(k string, v int) { *initial += v } },
		},
		{
			name:     "case 5",
			input:    nil,
			expected: 0,
			fn:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual int
			if tc.fn == nil {
				mapper.Each(tc.input, nil)
			} else {
				mapper.Each(tc.input, tc.fn(&actual))
			}

			assert.Equal(tt, tc.expected, actual)
		})
	}
}
