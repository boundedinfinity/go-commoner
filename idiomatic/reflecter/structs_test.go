package reflecter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/stretchr/testify/assert"
)

func Test_StructFieldNames(t *testing.T) {
	type Thing struct {
		Field1 string
		Field2 int
		Field3 bool
	}

	expected := []string{"Field1", "Field2", "Field3"}

	actual, err := reflecter.StructFieldNames[Thing]()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_StructFieldNamesAndTypes(t *testing.T) {
	type SubThing struct {
	}

	type Thing struct {
		Field1 string
		Field2 int
		Field3 bool
		Field4 SubThing
	}

	expected := map[string]string{
		"Field1": "string",
		"Field2": "int",
		"Field3": "bool",
		"Field4": "reflecter_test.SubThing",
	}

	actual, err := reflecter.StructFieldNamesAndTypes[Thing]()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
