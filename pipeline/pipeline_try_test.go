package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/commons/pipeline"
	"github.com/boundedinfinity/commons/strings"
	"github.com/boundedinfinity/commons/try"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Try(t *testing.T) {
	step1 := func(s string) try.Try[string] {
		return try.Success(strings.Capitalize(s))
	}

	excepted := "S"
	p := pipeline.NewTry[string]().Step(step1)
	actual := p.Run("s")

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}
