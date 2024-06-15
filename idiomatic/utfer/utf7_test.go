package utfer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
	"github.com/stretchr/testify/assert"
)

func Test_UTF7_IsUpperCase(t *testing.T) {
	assert.True(t, utfer.Utf7.IsUpperCase(byte('A')), "should be true")
	assert.False(t, utfer.Utf7.IsUpperCase(byte('a')), "should be false")
}

func Test_UTF7_IsLowerCase(t *testing.T) {
	assert.True(t, utfer.Utf7.IsLowerCase(byte('a')), "should be true")
	assert.False(t, utfer.Utf7.IsLowerCase(byte('A')), "should be false")
}

func Test_UTF7_IsSymbols(t *testing.T) {
	assert.True(t, utfer.Utf7.IsSymbol(byte('<')), "should be true")
	assert.False(t, utfer.Utf7.IsSymbol(byte('A')), "should be false")
}
