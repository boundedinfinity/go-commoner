package utf_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/utf"
	"github.com/stretchr/testify/assert"
)

func Test_IsUpperCase(t *testing.T) {
	assert.True(t, utf.Utf7.IsUpperCase(byte('A')))
	assert.False(t, utf.Utf7.IsUpperCase(byte('a')))
}

func Test_IsLowerCase(t *testing.T) {
	assert.True(t, utf.Utf7.IsLowerCase(byte('a')))
	assert.False(t, utf.Utf7.IsLowerCase(byte('A')))
}

func Test_IsSymbols(t *testing.T) {
	assert.True(t, utf.Utf7.IsSymbol(byte('<')))
	assert.False(t, utf.Utf7.IsSymbol(byte('A')))
}
