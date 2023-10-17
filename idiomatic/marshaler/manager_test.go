package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
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
        "name": "github.com/boundedinfinity/go-commoner/idiomatic/marshaler_test/aStruct",
        "instance": {
                "stuff": "some stuff"
            }
        }`
)

func Test_Marshal(t *testing.T) {
	input := theStruct
	expected := theJson

	actual, err := marshaler.MarshalIndent(input, "", "    ")
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func Test_Unmarshal(t *testing.T) {
	manager := marshaler.New()

	expected := theStruct
	input := theJson

	manager.Register(aStruct{})
	acutal, err := manager.Unmarshal([]byte(input))

	assert.Nil(t, err)
	assert.Equal(t, expected, acutal)
}

// func Test_UnmarshalTyped(t *testing.T) {
// 	expected := theStruct
// 	input := theJson

// 	marshaler.Register(expected)
// 	acutal, err := marshaler.UnmarshalTyped[aStruct]([]byte(input))

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected, acutal)
// }
