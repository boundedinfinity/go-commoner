package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Reject(t *testing.T) {
	type thing struct {
		a string
		b int
	}

	testCases := []struct {
		name     string
		input    []thing
		fn       func(i int, t thing) bool
		expected []thing
	}{
		{
			name:     "group by a",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: []thing{{a: "B", b: 30}},
		},
		{
			name:     "empty input",
			input:    []thing{},
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: []thing{},
		},
		{
			name:     "nil input",
			input:    nil,
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: []thing{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Reject(tc.fn, tc.input...)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}
