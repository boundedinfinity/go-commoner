package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Exist(t *testing.T) {
	type thing struct {
		a string
		b int
	}

	testCases := []struct {
		name     string
		input    []thing
		fn       func(i int, t thing) bool
		expected bool
	}{
		{
			name:     "should exist",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: true,
		},
		{
			name:     "should not exist",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       func(_ int, t thing) bool { return t.a == "Z" },
			expected: false,
		},
		{
			name:     "should exist empty input",
			input:    []thing{},
			fn:       func(_ int, t thing) bool { return t.a == "Z" },
			expected: false,
		},
		{
			name:     "should exist nil input",
			input:    nil,
			fn:       func(_ int, t thing) bool { return t.a == "A" },
			expected: false,
		},
		{
			name:     "should not exist with nil fn",
			input:    []thing{{a: "A", b: 10}, {a: "A", b: 20}, {a: "B", b: 30}},
			fn:       nil,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.ExistByI(tc.fn, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
