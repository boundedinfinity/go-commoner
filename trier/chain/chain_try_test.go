package chain_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/boundedinfinity/go-commoner/trier/chain"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Try(t *testing.T) {
	step1 := func(s string) string {
		return stringer.Capitalize(s)
	}

	excepted := "S"
	p := chain.New[string]().AppendNoError(step1)
	actual := p.Run("s")

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}
