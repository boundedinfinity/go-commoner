package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pipeline"
	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_with_string(t *testing.T) {
	excepted := "CBA"
	actual := pipeline.New[string]().
		Append(stringer.Uppercase[string]).
		Append(stringer.Reverse[string]).
		Append(stringer.TrimSpace[string]).
		RunSingle(" abc ")

	assert.Equal(t, excepted, actual)
}

func Test_Pipeline_with_slice(t *testing.T) {
	excepted := []string{"A", "B", "C"}
	input := []string{"a", "b", "c"}

	actual := pipeline.New[string]().
		Append(stringer.Uppercase[string]).
		Append(stringer.Reverse[string]).
		RunList(input)

	assert.Equal(t, excepted, actual)
}
