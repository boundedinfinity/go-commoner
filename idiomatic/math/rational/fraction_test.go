package rational_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/math/rational"
	"github.com/stretchr/testify/assert"
)

func Test_FractionComponent(t *testing.T) {
	assert.Equal(t, 14, rational.FractionComponent(3.14))
}

func Test_FractionSize(t *testing.T) {
	assert.Equal(t, 2, rational.FractionSize(3.14))
}

func Test_FractionFromFloat(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected rational.Fraction
	}{
		{
			name:  "1 / 32",
			input: float64(1) / float64(32),
			expected: rational.Fraction{
				Numerator:   1,
				Demoninator: 32,
			},
		},
		{
			name:  "2 / 32",
			input: float64(2) / float64(32),
			expected: rational.Fraction{
				Numerator:   1,
				Demoninator: 16,
			},
		},
		{
			name:  "3 / 32",
			input: float64(3) / float64(32),
			expected: rational.Fraction{
				Numerator:   3,
				Demoninator: 32,
			},
		},
		{
			name:  "4 / 32",
			input: float64(4) / float64(32),
			expected: rational.Fraction{
				Numerator:   1,
				Demoninator: 8,
			},
		},
		{
			name:  "5 / 32",
			input: float64(5) / float64(32),
			expected: rational.Fraction{
				Numerator:   5,
				Demoninator: 32,
			},
		},
		{
			name:  "6 / 32",
			input: float64(6) / float64(32),
			expected: rational.Fraction{
				Numerator:   3,
				Demoninator: 16,
			},
		},
		{
			name:  "7 / 32",
			input: float64(7) / float64(32),
			expected: rational.Fraction{
				Numerator:   7,
				Demoninator: 32,
			},
		},
		{
			name:  "8 / 32",
			input: float64(8) / float64(32),
			expected: rational.Fraction{
				Numerator:   1,
				Demoninator: 4,
			},
		},
		{
			name:  "9 / 32",
			input: float64(9) / float64(32),
			expected: rational.Fraction{
				Numerator:   9,
				Demoninator: 32,
			},
		},
		{
			name:  "10 / 32",
			input: float64(10) / float64(32),
			expected: rational.Fraction{
				Numerator:   5,
				Demoninator: 16,
			},
		},
		{
			name:  "11 / 32",
			input: float64(11) / float64(32),
			expected: rational.Fraction{
				Numerator:   11,
				Demoninator: 32,
			},
		},
		{
			name:  "12 / 32",
			input: float64(12) / float64(32),
			expected: rational.Fraction{
				Numerator:   3,
				Demoninator: 8,
			},
		},
		{
			name:  "13 / 32",
			input: float64(13) / float64(32),
			expected: rational.Fraction{
				Numerator:   13,
				Demoninator: 32,
			},
		},
		{
			name:  "14 / 32",
			input: float64(14) / float64(32),
			expected: rational.Fraction{
				Numerator:   7,
				Demoninator: 16,
			},
		},
		{
			name:  "15 / 32",
			input: float64(15) / float64(32),
			expected: rational.Fraction{
				Numerator:   15,
				Demoninator: 32,
			},
		},
		{
			name:  "16 / 32",
			input: float64(16) / float64(32),
			expected: rational.Fraction{
				Numerator:   1,
				Demoninator: 2,
			},
		},
		{
			name:  "17 / 32",
			input: float64(17) / float64(32),
			expected: rational.Fraction{
				Numerator:   17,
				Demoninator: 32,
			},
		},
		{
			name:  "18 / 32",
			input: float64(18) / float64(32),
			expected: rational.Fraction{
				Numerator:   9,
				Demoninator: 16,
			},
		},
		{
			name:  "19 / 32",
			input: float64(19) / float64(32),
			expected: rational.Fraction{
				Numerator:   19,
				Demoninator: 32,
			},
		},
		{
			name:  "20 / 32",
			input: float64(20) / float64(32),
			expected: rational.Fraction{
				Numerator:   5,
				Demoninator: 8,
			},
		},
		{
			name:  "21 / 32",
			input: float64(21) / float64(32),
			expected: rational.Fraction{
				Numerator:   21,
				Demoninator: 32,
			},
		},
		{
			name:  "22 / 32",
			input: float64(22) / float64(32),
			expected: rational.Fraction{
				Numerator:   11,
				Demoninator: 16,
			},
		},
		{
			name:  "23 / 32",
			input: float64(23) / float64(32),
			expected: rational.Fraction{
				Numerator:   23,
				Demoninator: 32,
			},
		},
		{
			name:  "24 / 32",
			input: float64(24) / float64(32),
			expected: rational.Fraction{
				Numerator:   3,
				Demoninator: 4,
			},
		},
		{
			name:  "25 / 32",
			input: float64(25) / float64(32),
			expected: rational.Fraction{
				Numerator:   25,
				Demoninator: 32,
			},
		},
		{
			name:  "26 / 32",
			input: float64(26) / float64(32),
			expected: rational.Fraction{
				Numerator:   13,
				Demoninator: 16,
			},
		},
		{
			name:  "27 / 32",
			input: float64(27) / float64(32),
			expected: rational.Fraction{
				Numerator:   27,
				Demoninator: 32,
			},
		},
		{
			name:  "28 / 32",
			input: float64(28) / float64(32),
			expected: rational.Fraction{
				Numerator:   7,
				Demoninator: 8,
			},
		},
		{
			name:  "29 / 32",
			input: float64(29) / float64(32),
			expected: rational.Fraction{
				Numerator:   29,
				Demoninator: 32,
			},
		},
		{
			name:  "30 / 32",
			input: float64(30) / float64(32),
			expected: rational.Fraction{
				Numerator:   15,
				Demoninator: 16,
			},
		},
		{
			name:  "31 / 32",
			input: float64(31) / float64(32),
			expected: rational.Fraction{
				Numerator:   31,
				Demoninator: 32,
			},
		},
		{
			name:  "32 / 32",
			input: float64(32) / float64(32),
			expected: rational.Fraction{
				Numerator:   0,
				Demoninator: 1,
			},
		},
		{
			name:  "3.14159",
			input: 3.14159,
			expected: rational.Fraction{
				Numerator:   14159,
				Demoninator: 100000,
			},
		},
		{
			name:  "2.71828",
			input: 2.71828,
			expected: rational.Fraction{
				Numerator:   17957,
				Demoninator: 25000,
			},
		},
		{
			name:  "3.14",
			input: 3.14,
			expected: rational.Fraction{
				Numerator:   7,
				Demoninator: 50,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := rational.FractionFromFloat(tc.input).Reduce()
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}
