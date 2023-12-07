package rational_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"
	"github.com/stretchr/testify/assert"
)

func Test_Rational_Whole(t *testing.T) {
	assert.Equal(t, 3, rational.Whole(3.14))
}

func Test_Rational_Reduce(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational
		expected rational.Rational
	}{
		{
			name:     "3.75",
			input:    rational.New(3, 75, 100),
			expected: rational.New(3, 3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Reduce()
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_Rational_Mixed(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational
		expected rational.Rational
	}{
		{
			name:     "14/3",
			input:    rational.New(0, 14, 3),
			expected: rational.New(4, 2, 3),
		},
		{
			name:     "2 14/3",
			input:    rational.New(2, 14, 3),
			expected: rational.New(6, 2, 3),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Mixed()
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_Rational_Improper(t *testing.T) {
	testCases := []struct {
		name     string
		input    rational.Rational
		expected rational.Rational
	}{
		{
			name:     "2 8/32",
			input:    rational.New(2, 8, 32).Reduce(),
			expected: rational.New(0, 9, 4),
		},
		{
			name:     "1 3/4",
			input:    rational.New(1, 3, 4),
			expected: rational.New(0, 7, 4),
		},
		{
			name:     "3/4",
			input:    rational.New(0, 3, 4),
			expected: rational.New(0, 3, 4),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Improper()
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}
