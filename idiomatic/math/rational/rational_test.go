package rational_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/math/rational"
	"github.com/stretchr/testify/assert"
)

func Test_IntegerComponent(t *testing.T) {
	assert.Equal(t, 3, rational.IntegerComponent(3.14))
}
