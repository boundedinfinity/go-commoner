package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pipeline"
	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline(t *testing.T) {
	step1 := func(s string) (string, error) {
		return strings.Capitalize(s), nil
	}

	excepted := "S"
	p := pipeline.New[string]().Step(step1)
	actual, err := p.Run("s")

	assert.Nil(t, err)
	assert.Equal(t, excepted, actual)
}
