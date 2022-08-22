package strings_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_IsLetters(t *testing.T) {
	assert.False(t, strings.IsLetters("0"))
	assert.True(t, strings.IsLetters("a"))
}

func Test_IsNumbers(t *testing.T) {
	assert.True(t, strings.IsNumbers("0"))
	assert.False(t, strings.IsNumbers("a"))
}
