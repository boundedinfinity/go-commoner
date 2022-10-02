package utf8_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/utf8"
	"github.com/stretchr/testify/assert"
)

func Test_IsUpperCase(t *testing.T) {
	assert.True(t, utf8.IsUpperCase(byte('A')))
	assert.False(t, utf8.IsUpperCase(byte('a')))
}

func Test_IsLowerCase(t *testing.T) {
	assert.True(t, utf8.IsLowerCase(byte('a')))
	assert.False(t, utf8.IsLowerCase(byte('A')))
}

func Test_IsSymbols(t *testing.T) {
	assert.True(t, utf8.IsSymbols(byte('<')))
	assert.False(t, utf8.IsSymbols(byte('A')))
}
