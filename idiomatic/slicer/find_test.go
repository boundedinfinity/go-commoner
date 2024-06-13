package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Find(t *testing.T) {
	type thing struct {
		a string
		b int
	}

	testCases := []struct {
		name     string
		input    []thing
		fn       func(i int, t thing) bool
		expected thing
		ok       bool
	}{
		{
			name:     "should find",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: thing{a: "A", b: 10},
			ok:       true,
		},
		{
			name:     "should not find",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       func(_ int, t thing) bool { return t.a == "Z" },
			expected: thing{},
			ok:       false,
		},
		{
			name:     "should not find with empty input",
			input:    []thing{},
			fn:       func(_ int, t thing) bool { return t.a == "Z" },
			expected: thing{},
			ok:       false,
		},
		{
			name:     "should not find with nil input",
			input:    nil,
			fn:       func(_ int, t thing) bool { return t.a == "Z" },
			expected: thing{},
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, ok := slicer.FindFn(tc.fn, tc.input...)
			assert.Equal(tt, tc.expected, actual)
			assert.Equal(tt, tc.ok, ok)
		})
	}
}
