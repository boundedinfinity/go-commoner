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

type thingA struct {
	ThingA string
}

type thingB struct {
	ThingB string
}

type someInterface interface {
	GetThing() string
}

func Test_Marshaler(t *testing.T) {
	m := marshaler.New()

	var ta thingA
	var tb thingB

	m.Interface(&WhichThing{})
	m.Register(ta, func() any { return thingA{} })
	m.Register(tb, func() any { return thingB{} })

	actual1, err := m.UnmarshalInterface([]byte(`
		{
			"Type": "thingA",
			"ThingA": "ThingA"
		}
	`))

	assert.Nil(t, err)
	assert.Equal(t, ta, actual1)
}
