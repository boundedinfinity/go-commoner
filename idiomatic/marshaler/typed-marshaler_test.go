package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/stretchr/testify/assert"
)

type typedMarshalerThing struct {
	Type string
}

func (t typedMarshalerThing) Discriminator() string {
	return t.Type
}

func (t typedMarshalerThing) Value() any {
	return t
}

type typedThingA struct {
	ThingA string
}

type typedThingB struct {
	ThingB string
}

type typedInterface interface {
	GetThing() string
}

var (
	typedMessage = `
    {
        "type": "github.com/boundedinfinity/go-commoner/idiomatic/marshaler_test/typedThingA",
        "value": {
            "ThingA": "somethingA"
        }
    }
`
)

func Test_TypedMarshaler_Unmarshal(t *testing.T) {
	m := marshaler.NewTyped()
	m.Register(typedThingA{})
	m.Register(typedThingB{})

	actual1, err := m.Unmarshal([]byte(typedMessage))
	assert.Nil(t, err)

	actualThing, ok := actual1.(typedThingA)
	assert.True(t, ok)

	assert.Equal(t, typedThingA{"somethingA"}, actualThing)
}

func Test_TypedMarshaler_Marshal(t *testing.T) {
	m := marshaler.NewTyped()
	m.Register(typedThingA{})
	bs, err := m.Marshal(typedThingA{ThingA: "somethingA"})

	assert.Nil(t, err)

	actual := string(bs)
	assert.JSONEq(t, typedMessage, actual)
}
