package strings_test

import (
	"testing"

	"github.com/boundedinfinity/commons/strings"
	"github.com/stretchr/testify/assert"
)

func Test_String_IsBlank(t *testing.T) {
	assert.Equal(t, strings.IsBlank(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, strings.IsBlank(ms), true)
}

func Test_String_IsEmpty(t *testing.T) {
	assert.Equal(t, strings.IsEmpty(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, strings.IsBlank(ms), true)
}
