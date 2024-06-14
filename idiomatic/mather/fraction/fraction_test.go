package fraction_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
	"github.com/stretchr/testify/assert"
)

func Test_Fraction_String(t *testing.T) {
	testCases := []struct {
		name     string
		input    fraction.Fraction
		expected string
	}{
		{
			name:     "1/2",
			input:    fraction.New(1, 2),
			expected: "1/2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.String()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_Component(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "1/2",
			input:    1.0 / 2.0,
			expected: 5,
		},
		{
			name:     "3.14",
			input:    3.14,
			expected: 14,
		},
		{
			name:     "0.14",
			input:    0.14,
			expected: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.Component(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_Magnitude(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "1/2",
			input:    1.0 / 2.0,
			expected: 1,
		},
		{
			name:     "3.14",
			input:    3.14,
			expected: 2,
		},
		{
			name:     "0.14",
			input:    0.14,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.Magnitude(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fraction_FromFloat(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected fraction.Fraction
	}{
		{
			name:  "1 / 32",
			input: float64(1) / float64(32),
			expected: fraction.Fraction{
				Numerator:   1,
				Denominator: 32,
			},
		},
		{
			name:  "2 / 32",
			input: float64(2) / float64(32),
			expected: fraction.Fraction{
				Numerator:   1,
				Denominator: 16,
			},
		},
		{
			name:  "3 / 32",
			input: float64(3) / float64(32),
			expected: fraction.Fraction{
				Numerator:   3,
				Denominator: 32,
			},
		},
		{
			name:  "4 / 32",
			input: float64(4) / float64(32),
			expected: fraction.Fraction{
				Numerator:   1,
				Denominator: 8,
			},
		},
		{
			name:  "5 / 32",
			input: float64(5) / float64(32),
			expected: fraction.Fraction{
				Numerator:   5,
				Denominator: 32,
			},
		},
		{
			name:  "6 / 32",
			input: float64(6) / float64(32),
			expected: fraction.Fraction{
				Numerator:   3,
				Denominator: 16,
			},
		},
		{
			name:  "7 / 32",
			input: float64(7) / float64(32),
			expected: fraction.Fraction{
				Numerator:   7,
				Denominator: 32,
			},
		},
		{
			name:  "8 / 32",
			input: float64(8) / float64(32),
			expected: fraction.Fraction{
				Numerator:   1,
				Denominator: 4,
			},
		},
		{
			name:  "9 / 32",
			input: float64(9) / float64(32),
			expected: fraction.Fraction{
				Numerator:   9,
				Denominator: 32,
			},
		},
		{
			name:  "10 / 32",
			input: float64(10) / float64(32),
			expected: fraction.Fraction{
				Numerator:   5,
				Denominator: 16,
			},
		},
		{
			name:  "11 / 32",
			input: float64(11) / float64(32),
			expected: fraction.Fraction{
				Numerator:   11,
				Denominator: 32,
			},
		},
		{
			name:  "12 / 32",
			input: float64(12) / float64(32),
			expected: fraction.Fraction{
				Numerator:   3,
				Denominator: 8,
			},
		},
		{
			name:  "13 / 32",
			input: float64(13) / float64(32),
			expected: fraction.Fraction{
				Numerator:   13,
				Denominator: 32,
			},
		},
		{
			name:  "14 / 32",
			input: float64(14) / float64(32),
			expected: fraction.Fraction{
				Numerator:   7,
				Denominator: 16,
			},
		},
		{
			name:  "15 / 32",
			input: float64(15) / float64(32),
			expected: fraction.Fraction{
				Numerator:   15,
				Denominator: 32,
			},
		},
		{
			name:  "16 / 32",
			input: float64(16) / float64(32),
			expected: fraction.Fraction{
				Numerator:   1,
				Denominator: 2,
			},
		},
		{
			name:  "17 / 32",
			input: float64(17) / float64(32),
			expected: fraction.Fraction{
				Numerator:   17,
				Denominator: 32,
			},
		},
		{
			name:  "18 / 32",
			input: float64(18) / float64(32),
			expected: fraction.Fraction{
				Numerator:   9,
				Denominator: 16,
			},
		},
		{
			name:  "19 / 32",
			input: float64(19) / float64(32),
			expected: fraction.Fraction{
				Numerator:   19,
				Denominator: 32,
			},
		},
		{
			name:  "20 / 32",
			input: float64(20) / float64(32),
			expected: fraction.Fraction{
				Numerator:   5,
				Denominator: 8,
			},
		},
		{
			name:  "21 / 32",
			input: float64(21) / float64(32),
			expected: fraction.Fraction{
				Numerator:   21,
				Denominator: 32,
			},
		},
		{
			name:  "22 / 32",
			input: float64(22) / float64(32),
			expected: fraction.Fraction{
				Numerator:   11,
				Denominator: 16,
			},
		},
		{
			name:  "23 / 32",
			input: float64(23) / float64(32),
			expected: fraction.Fraction{
				Numerator:   23,
				Denominator: 32,
			},
		},
		{
			name:  "24 / 32",
			input: float64(24) / float64(32),
			expected: fraction.Fraction{
				Numerator:   3,
				Denominator: 4,
			},
		},
		{
			name:  "25 / 32",
			input: float64(25) / float64(32),
			expected: fraction.Fraction{
				Numerator:   25,
				Denominator: 32,
			},
		},
		{
			name:  "26 / 32",
			input: float64(26) / float64(32),
			expected: fraction.Fraction{
				Numerator:   13,
				Denominator: 16,
			},
		},
		{
			name:  "27 / 32",
			input: float64(27) / float64(32),
			expected: fraction.Fraction{
				Numerator:   27,
				Denominator: 32,
			},
		},
		{
			name:  "28 / 32",
			input: float64(28) / float64(32),
			expected: fraction.Fraction{
				Numerator:   7,
				Denominator: 8,
			},
		},
		{
			name:  "29 / 32",
			input: float64(29) / float64(32),
			expected: fraction.Fraction{
				Numerator:   29,
				Denominator: 32,
			},
		},
		{
			name:  "30 / 32",
			input: float64(30) / float64(32),
			expected: fraction.Fraction{
				Numerator:   15,
				Denominator: 16,
			},
		},
		{
			name:  "31 / 32",
			input: float64(31) / float64(32),
			expected: fraction.Fraction{
				Numerator:   31,
				Denominator: 32,
			},
		},
		{
			name:  "32 / 32",
			input: float64(32) / float64(32),
			expected: fraction.Fraction{
				Numerator:   0,
				Denominator: 1,
			},
		},
		{
			name:  "3.14159",
			input: 3.14159,
			expected: fraction.Fraction{
				Numerator:   14159,
				Denominator: 100000,
			},
		},
		{
			name:  "2.71828",
			input: 2.71828,
			expected: fraction.Fraction{
				Numerator:   17957,
				Denominator: 25000,
			},
		},
		{
			name:  "3.14",
			input: 3.14,
			expected: fraction.Fraction{
				Numerator:   7,
				Denominator: 50,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := fraction.FromFloat(tc.input).Reduce()
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
