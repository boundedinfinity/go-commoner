package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pipeline"
	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline_Try(t *testing.T) {
	step1 := pipeline.SimpleStep(func(s string) string {
		return stringer.Capitalize(s)
	})

	excepted := "S"
	p := pipeline.NewTry[string]().Append(step1)
	actual := p.Run("s")

	assert.False(t, actual.Failure())
	assert.True(t, actual.Success())
	assert.Equal(t, excepted, actual.Result)
}
