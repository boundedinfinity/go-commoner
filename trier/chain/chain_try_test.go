package chain_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/boundedinfinity/go-commoner/trier/chain"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Try(t *testing.T) {
	excepted := "S"
	p := chain.New[string]().AppendNoError(stringer.Capitalize[string])
	actual := p.RunSingle("s")

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}
