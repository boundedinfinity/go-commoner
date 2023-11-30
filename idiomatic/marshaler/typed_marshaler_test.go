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

type wrappedThingA struct {
	ThingA string
}

type wrappedThingB struct {
	ThingB string
}

type someInterface interface {
	GetThing() string
}

var (
	typedMessage = `
    {
        "type": "wrappedThingA",
        "value": {
            "ThingA": "somethingA"
        }
    }
`
)

func Test_TypedMarshaler_Unmarshal(t *testing.T) {
	m := marshaler.NewTyped()
	m.Register(wrappedThingA{})
	m.Register(wrappedThingB{})

	actual1, err := m.Unmarshal([]byte(typedMessage))
	assert.Nil(t, err)

	actualThing, ok := actual1.(wrappedThingA)
	assert.True(t, ok)

	assert.Equal(t, wrappedThingA{"somethingA"}, actualThing)
}

func Test_TypedMarshaler_Marshal(t *testing.T) {
	m := marshaler.NewTyped()
	m.Register(wrappedThingA{})
	bs, err := m.Marshal(wrappedThingA{ThingA: "somethingA"})

	assert.Nil(t, err)

	actual := string(bs)
	assert.JSONEq(t, typedMessage, actual)
}
