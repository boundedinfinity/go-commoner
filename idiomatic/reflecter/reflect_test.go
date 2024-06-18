package reflecter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/stretchr/testify/assert"
)

type aString string

func Test_Instances_QualifiedName(t *testing.T) {
	actual := reflecter.TypeQualifiedName[aString]()
	assert.Equal(t, "github.com/boundedinfinity/go-commoner/idiomatic/reflecter_test/aString", actual)
}

func Test_Instances_BaseName(t *testing.T) {
	actual := reflecter.TypeBaseName[aString]()
	assert.Equal(t, "aString", actual)
}
