package utf_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/utf"
	"github.com/stretchr/testify/assert"
)

func Test_HtmlNumber(t *testing.T) {
	assert.Equal(t, "&#00;", utf.NULL.HtmlNumber())
	assert.Equal(t, "&#65;", utf.UPPERCASE_A.HtmlNumber())
}
