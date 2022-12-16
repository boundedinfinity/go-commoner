package chain_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/chain"
	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_with_single(t *testing.T) {
	excepted := "CBA"
	actual, err := chain.
		Then(stringer.Uppercase[string]).
		Then(stringer.Reverse[string]).
		Then(stringer.TrimSpace[string]).
		Run(" abc ")

	assert.Nil(t, err)
	assert.Equal(t, excepted, actual)
}

func Test_Pipeline_with_each(t *testing.T) {
	excepted := []chain.ChainItem[string]{{"A", nil}, {"B", nil}, {"C", nil}}
	input := []string{"a", "b", "c"}

	actual := chain.
		Then(stringer.Uppercase[string]).
		Then(stringer.Reverse[string]).
		Each(input)

	assert.Equal(t, len(excepted), len(actual))
	assert.Equal(t, excepted[0].Result, actual[0].Result)
}
