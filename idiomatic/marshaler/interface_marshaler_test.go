package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/stretchr/testify/assert"
)

type interfaceThing struct {
	Field1 string
}

type interfaceThingDiscriminator struct {
	Kind string
	interfaceThing
}

func (t interfaceThingDiscriminator) Discriminator() string {
	return t.Kind
}

var _ marshaler.InterfaceUnmarshalDiscriminator = &interfaceThingDiscriminator{}

var (
	interfaceMessage = `
    {
        "kind": "kindA",
        "Field1": "someField1"
    }
`
)

func Test_InterfaceMarshaler_Unmarshal(t *testing.T) {
	m := marshaler.NewInterface(&interfaceThingDiscriminator{})
	// m.Register(wrappedThingA{})
	// m.Register(wrappedThingB{})

	// actual1, err := m.Unmarshal([]byte(interfaceMessage))
	// assert.Nil(t, err)

	// actualThing, ok := actual1.(wrappedThingA)
	// assert.True(t, ok)

	// assert.Equal(t, wrappedThingA{"somethingA"}, actualThing)
}

func Test_InterfaceMarshaler_Marshal(t *testing.T) {
	m := marshaler.NewWrapped()
	m.Register(wrappedThingA{})
	bs, err := m.Marshal(wrappedThingA{ThingA: "somethingA"})

	assert.Nil(t, err)

	actual := string(bs)
	assert.JSONEq(t, interfaceMessage, actual)
}
