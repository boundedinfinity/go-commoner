package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Group(t *testing.T) {
	type thing struct {
		a string
		b int
	}

	testCases := []struct {
		name     string
		input    []thing
		fn       func(i int, t thing) string
		expected map[string][]thing
	}{
		{
			name:  "group by a",
			input: []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:    func(_ int, t thing) string { return t.a },
			expected: map[string][]thing{
				"A": {{a: "A", b: 10}, {a: "A", b: 20}},
				"B": {{a: "B", b: 30}},
			},
		},
		{
			name:     "empty input",
			input:    []thing{},
			fn:       func(_ int, t thing) string { return t.a },
			expected: map[string][]thing{},
		},
		{
			name:     "nil input",
			input:    nil,
			fn:       func(_ int, t thing) string { return t.a },
			expected: map[string][]thing{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Group(tc.fn, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}
