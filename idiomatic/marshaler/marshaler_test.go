package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/stretchr/testify/assert"
)

type WhichThing struct {
	Type string
}

func (t WhichThing) Discriminator() string {
	return t.Type
}

func (t WhichThing) Value() any {
	return t
}

type thingA struct {
	ThingA string
}

type thingB struct {
	ThingB string
}

type someInterface interface {
	GetThing() string
}

var (
	message = `
    {
        "type": "thingA",
        "value": {
            "ThingA": "somethingA"
        }
    }
`
)

func Test_Marshaler_Unmarshal(t *testing.T) {
	m := marshaler.New()
	m.Register(thingA{})
	m.Register(thingB{})

	actual1, err := m.Unmarshal([]byte(message))
	assert.Nil(t, err)

	actualThing, ok := actual1.(thingA)
	assert.True(t, ok)

	assert.Equal(t, thingA{"somethingA"}, actualThing)
}

func Test_Marshaler_Marshal(t *testing.T) {
	m := marshaler.New()
	m.Register(thingA{})
	bs, err := m.Marshal(thingA{ThingA: "somethingA"})

	assert.Nil(t, err)

	actual := string(bs)
	assert.JSONEq(t, message, actual)
}
