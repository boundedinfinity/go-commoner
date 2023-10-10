package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Scrub(t *testing.T) {
	testCases := []struct {
		name        string
		fn          func(s string, r string) string
		replacement string
		input       string
		expected    string
	}{
		{
			name:        "replace numbers with _",
			fn:          stringer.ReplaceNumbers[string],
			replacement: "_",
			input:       "x 1 x 2 x 3",
			expected:    "x _ x _ x _",
		},
		{
			name:        "replace letters with _",
			fn:          stringer.ReplaceLetters[string],
			replacement: "_",
			input:       "x 1 x 2 x 3",
			expected:    "_ 1 _ 2 _ 3",
		},
		{
			name:        "replace symbols with _",
			fn:          stringer.ReplaceSymbols[string],
			replacement: "_",
			input:       "x ~ x ! x @",
			expected:    "x _ x _ x _",
		},
		{
			name:        "replace non-words characters with _",
			fn:          stringer.ReplaceSymbols[string],
			replacement: "_",
			input:       "x ~ 9 ! x @",
			expected:    "x _ 9 _ x _",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input, tc.replacement)
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}
