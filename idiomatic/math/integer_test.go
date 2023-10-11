package math_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/stretchr/testify/assert"
)

func Test_IntegerComponent(t *testing.T) {
	assert.Equal(t, 3, math.GreatestCommonDivisor(3, 5))
}
