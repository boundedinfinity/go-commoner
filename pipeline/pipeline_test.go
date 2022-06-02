package pipeline_test

import (
	"testing"

	"github.com/boundedinfinity/commons/pipeline"
	"github.com/stretchr/testify/assert"
)

func Test_Pipeline(t *testing.T) {
	p := pipeline.New()

	assert.NotNil(t, p)
}
