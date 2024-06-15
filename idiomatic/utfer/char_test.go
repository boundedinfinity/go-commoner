package utfer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/utfer"
	"github.com/stretchr/testify/assert"
)

func Test_HtmlNumber(t *testing.T) {
	assert.Equal(t, "&#00;", utfer.NULL.HtmlNumber())
	assert.Equal(t, "&#65;", utfer.UPPERCASE_A.HtmlNumber())
}
