package chain_test

import (
	"path/filepath"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-commoner/trier/chain"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Then(t *testing.T) {
	input := "something"
	excepted := "Something"
	actual := chain.
		Then(stringer.Capitalize[string]).
		Run(input)

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}

func Test_Pipeline_WithErr(t *testing.T) {
	input := "../something"
	// excepted := "Something"
	actual := chain.
		CanErr(filepath.Abs).
		Run(input)

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	// assert.Equal(t, excepted, actual.Result)
}
