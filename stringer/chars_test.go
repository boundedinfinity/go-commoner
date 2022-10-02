package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_IsLetters(t *testing.T) {
	assert.False(t, stringer.IsLetters("0"))
	assert.True(t, stringer.IsLetters("a"))
}

func Test_IsNumbers(t *testing.T) {
	assert.True(t, stringer.IsNumbers("0"))
	assert.False(t, stringer.IsNumbers("a"))
}
