package caser_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/caser"
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

func Test_Phrase_to_Camel(t *testing.T) {
	actual := caser.PhraseToCamel(aPhrase)
	assert.Equal(t, camel, actual)
}

func Test_Phrase_to_Pascal(t *testing.T) {
	actual := caser.PhraseToPascal(aPhrase)
	assert.Equal(t, pascal, actual)
}

func Test_Phrase_to_Snake_Lower(t *testing.T) {
	actual := caser.PhraseToSnake(aPhrase)
	assert.Equal(t, snakeLower, actual)
}

func Test_Phrase_to_Snake_Upper(t *testing.T) {
	actual := caser.PhraseToSnakeUpper(aPhrase)
	assert.Equal(t, snakeUpper, actual)
}

func Test_Phrase_to_Kabab_Lower(t *testing.T) {
	actual := caser.PhraseToKebab(aPhrase)
	assert.Equal(t, kababLower, actual)
}

func Test_Phrase_to_Kabab_Upper(t *testing.T) {
	actual := caser.PhraseToKebabUpper(aPhrase)
	assert.Equal(t, kababUpper, actual)
}

func Test_Camel_to_Phrase(t *testing.T) {
	actual := caser.CamelToPhrase(camel)
	assert.Equal(t, aPhrase, actual)
}

func Test_Pascal_to_Phrase(t *testing.T) {
	actual := caser.CamelToPhrase(pascal)
	assert.Equal(t, aPhrase, actual)
}

func Test_Snake_Lower_to_Phrase(t *testing.T) {
	actual := caser.SnakeToPhrase(snakeLower)
	assert.Equal(t, aPhrase, actual)
}

func Test_Snake_Upper_to_Phrase(t *testing.T) {
	actual := caser.SnakeToPhrase(snakeUpper)
	assert.Equal(t, aPhrase, actual)
}

func Test_Kabab_Lower_to_Phrase(t *testing.T) {
	actual := caser.KebabToPhrase(kababLower)
	assert.Equal(t, aPhrase, actual)
}

func Test_Kabab_Upper_to_Phrase(t *testing.T) {
	actual := caser.KebabToPhrase(kababUpper)
	assert.Equal(t, aPhrase, actual)
}
