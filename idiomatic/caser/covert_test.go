package caser_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/stretchr/testify/assert"
)

var (
	aPhrase    = "Some Test 4 Thing"
	camel      = "someTest4Thing"
	pascal     = "SomeTest4Thing"
	snakeLower = "some_test_4_thing"
	snakeUpper = "SOME_TEST_4_THING"
	kababLower = "some-test-4-thing"
	kababUpper = "SOME-TEST-4-THING"
)

func Test_Convert(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
		fn       func(string) string
	}{
		{
			name:     "Phrase_to_Camel",
			input:    aPhrase,
			expected: camel,
			fn:       caser.PhraseToCamel[string],
		},
		{
			name:     "Phrase_to_Pascal",
			input:    aPhrase,
			expected: pascal,
			fn:       caser.PhraseToPascal[string],
		},
		{
			name:     "Phrase_to_Snake_Lower",
			input:    aPhrase,
			expected: snakeLower,
			fn:       caser.PhraseToSnake[string],
		},
		{
			name:     "Phrase_to_Snake_Upper",
			input:    aPhrase,
			expected: snakeUpper,
			fn:       caser.PhraseToSnakeUpper[string],
		},
		{
			name:     "Phrase_to_Kabab_Lower",
			input:    aPhrase,
			expected: kababLower,
			fn:       caser.PhraseToKebab[string],
		},
		{
			name:     "Phrase_to_Kabab_Upper",
			input:    aPhrase,
			expected: kababUpper,
			fn:       caser.PhraseToKebabUpper[string],
		},
		{
			name:     "Camel_to_Phrase",
			input:    camel,
			expected: aPhrase,
			fn:       caser.CamelToPhrase[string],
		},
		{
			name:     "Pascal_to_Phrase",
			input:    pascal,
			expected: aPhrase,
			fn:       caser.PascalToPhrase[string],
		},
		{
			name:     "Snake_Lower_to_Phrase",
			input:    snakeLower,
			expected: aPhrase,
			fn:       caser.SnakeToPhrase[string],
		},
		{
			name:     "Snake_Upper_to_Phrase",
			input:    snakeUpper,
			expected: aPhrase,
			fn:       caser.SnakeToPhrase[string],
		},
		{
			name:     "Kabab_Lower_to_Phrase",
			input:    kababLower,
			expected: aPhrase,
			fn:       caser.KebabToPhrase[string],
		},
		{
			name:     "Kabab_Upper_to_Phrase",
			input:    kababUpper,
			expected: aPhrase,
			fn:       caser.KebabToPhrase[string],
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.fn(tc.input)
			assert.Equal(t, actual, tc.expected)
		})
	}
}

func Test_ParseConvert(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
		convert  string
		err      error
	}{
		{
			name:     "Phrase_to_Camel",
			input:    aPhrase,
			expected: camel,
			convert:  "phrase-to-camel",
			err:      nil,
		},
		{
			name:     "Phrase_to_Pascal",
			input:    aPhrase,
			expected: pascal,
			convert:  "phrase-to-pascal",
			err:      nil,
		},
		{
			name:     "Phrase_to_Snake",
			input:    aPhrase,
			expected: snakeLower,
			convert:  "phrase-to-snake",
			err:      nil,
		},
		{
			name:     "Phrase_to_Snake_Lower",
			input:    aPhrase,
			expected: snakeLower,
			convert:  "phrase-to-snake-lower",
			err:      nil,
		},
		{
			name:     "Phrase_to_Snake_Upper",
			input:    aPhrase,
			expected: snakeUpper,
			convert:  "phrase-to-snake-upper",
			err:      nil,
		},
		{
			name:     "Phrase_to_Kabab_Lower",
			input:    aPhrase,
			expected: kababLower,
			convert:  "phrase-to-kebab",
			err:      nil,
		},
		{
			name:     "Phrase_to_Kabab_Upper",
			input:    aPhrase,
			expected: kababUpper,
			convert:  "phrase-to-kebab-upper",
			err:      nil,
		},
		{
			name:     "Camel_to_Phrase",
			input:    camel,
			expected: aPhrase,
			convert:  "camel-to-phrase",
			err:      nil,
		},
		{
			name:     "Pascal_to_Phrase",
			input:    pascal,
			expected: aPhrase,
			convert:  "pascal-to-phrase",
			err:      nil,
		},
		{
			name:     "Snake_Lower_to_Phrase",
			input:    snakeLower,
			expected: aPhrase,
			convert:  "snake-to-phrase",
			err:      nil,
		},
		{
			name:     "Snake_Upper_to_Phrase",
			input:    snakeUpper,
			expected: aPhrase,
			convert:  "snake-upper-to-phrase",
			err:      nil,
		},
		{
			name:     "Kabab_to_Phrase",
			input:    kababLower,
			expected: aPhrase,
			convert:  "kebab-to-phrase",
			err:      nil,
		},
		{
			name:     "Kabab_Lower_to_Phrase",
			input:    kababLower,
			expected: aPhrase,
			convert:  "kebab-lower-to-phrase",
			err:      nil,
		},
		{
			name:     "Kabab_Upper_to_Phrase",
			input:    kababUpper,
			expected: aPhrase,
			convert:  "kebab-upper-to-phrase",
			err:      nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := caser.ConvertAs(tc.input, tc.convert)

			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, tc.err, err)
			}

			assert.Equal(t, tc.err, err)
			assert.Equal(t, actual, tc.expected, tc.convert)
		})
	}
}
