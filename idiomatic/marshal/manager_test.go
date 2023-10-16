package marshal_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshal"
	"github.com/stretchr/testify/assert"
)

type aStruct struct {
	Stuff string `json:"stuff,omitempty"`
}

var (
	theStruct = aStruct{
		Stuff: "some stuff",
	}

	theJson = `{
        "name": "github.com/boundedinfinity/go-commoner/idiomatic/marshal_test/aStruct",
        "instance": {
                "stuff": "some stuff"
            }
        }`
)

func Test_Marshal(t *testing.T) {
	input := theStruct
	expected := theJson

	actual, err := marshal.MarshalIndent(input, "", "    ")
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func Test_Unmarshal(t *testing.T) {
	expected := theStruct
	input := theJson

	marshal.Register(expected)
	acutal, err := marshal.Unmarshal([]byte(input))

	assert.Nil(t, err)
	assert.Equal(t, &expected, acutal)
}

func Test_UnmarshalTyped(t *testing.T) {
	expected := theStruct
	input := theJson

	marshal.Register(expected)
	acutal, err := marshal.UnmarshalTyped[aStruct]([]byte(input))

	assert.Nil(t, err)
	assert.Equal(t, expected, acutal)
}
