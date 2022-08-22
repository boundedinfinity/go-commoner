package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pipeline"
	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/boundedinfinity/go-trier"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Try(t *testing.T) {
	step1 := func(s string) trier.Try[string] {
		return trier.Success(strings.Capitalize(s))
	}

	excepted := "S"
	p := pipeline.NewTry[string]().Step(step1)
	actual := p.Run("s")

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}
