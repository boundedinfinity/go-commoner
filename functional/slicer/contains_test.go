package slicer_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/slicer"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	testCases := []struct {
		name     string
		match    string
		input    []string
		expected bool
	}{
		{
			name:     "contains a",
			match:    "a",
			input:    []string{"b", "c", "b", "a"},
			expected: true,
		},
		{
			name:     "does not contain a",
			match:    "a",
			input:    []string{"b", "c", "d", "e"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Contains(tc.match, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_ContainsFnErr(t *testing.T) {
	testCases := []struct {
		name     string
		match    func(int, string) (bool, error)
		input    []string
		expected trier.Try[bool]
	}{
		{
			name:     "contains a",
			match:    func(_ int, s string) (bool, error) { return s == "a", nil },
			input:    []string{"b", "c", "b", "a"},
			expected: trier.Success(true),
		},
		{
			name:     "does not contain a",
			match:    func(_ int, s string) (bool, error) { return s == "a", nil },
			input:    []string{"b", "c", "d", "e"},
			expected: trier.Success(false),
		},
		{
			name:     "does not contain a",
			match:    func(_ int, s string) (bool, error) { return false, errors.New("failed") },
			input:    []string{"b", "c", "d", "e"},
			expected: trier.Failure[bool](errors.New("failed")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.ContainsByErrI(tc.match, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
