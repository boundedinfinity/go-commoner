package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_String_IsZero(t *testing.T) {
	assert.Equal(t, stringer.IsZero(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, stringer.IsZero(ms), true)
}

func Test_String_FindNonZero(t *testing.T) {
	actual1, ok1 := stringer.FindNonZero("", "x")
	assert.Equal(t, "x", actual1)
	assert.Equal(t, true, ok1)

	type MyString string
	zero := MyString("")
	x := MyString("x")
	actual2, ok2 := stringer.FindNonZero(zero, MyString(x))
	assert.Equal(t, x, actual2)
	assert.Equal(t, true, ok2)
}
