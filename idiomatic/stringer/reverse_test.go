package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Stringer_Reverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "a b c d",
			expected: "d c b a",
		},
		{
			name:     "case 2",
			input:    "",
			expected: "",
		},
		{
			name:     "case 3",
			input:    "a",
			expected: "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.Reverse(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Benchmark_Stringer_Reverse(b *testing.B) {
	for b.Loop() {
		stringer.Reverse("a b c d")
	}
}
