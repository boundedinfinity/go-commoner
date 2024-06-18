package utfer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
	"github.com/stretchr/testify/assert"
)

func Test_Scrub(t *testing.T) {
	testCases := []struct {
		name     string
		fn       func(string, byte) string
		new      byte
		input    string
		expected string
	}{
		{
			name:     "replace numbers with _",
			fn:       utfer.ReplaceNumbers[string],
			new:      '_',
			input:    "x 1 x 2 x 3",
			expected: "x _ x _ x _",
		},
		{
			name:     "replace letters with _",
			fn:       utfer.ReplaceLetters[string],
			new:      '_',
			input:    "x 1 x 2 x 3",
			expected: "_ 1 _ 2 _ 3",
		},
		{
			name:     "replace symbols with _",
			fn:       utfer.ReplaceSymbols[string],
			new:      '_',
			input:    "x ~ x ! x @",
			expected: "x _ x _ x _",
		},
		{
			name:     "replace non-words characters with _",
			fn:       utfer.ReplaceSymbols[string],
			new:      '_',
			input:    "x ~ 9 ! x @",
			expected: "x _ 9 _ x _",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input, byte(tc.new))
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}
