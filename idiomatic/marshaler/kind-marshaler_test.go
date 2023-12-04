package marshaler_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/stretchr/testify/assert"
)

type Animal struct {
	Kind string
}

type Dog struct {
	Kind string
	Bark string
}

type Cat struct {
	Kind string
	Meow string
}

var (
	dogJson = `
    {
        "Kind": "dog",
        "Bark": "load"
    }
`
	catJson = `
    {
        "Kind": "cat",
        "Meow": "low"
    }
`
)

func Test_InterfaceMarshaler_Unmarshal(t *testing.T) {
	m := marshaler.NewKind[string, Animal]()

	var dog Dog
	var cat Cat
	var err error

	m.RegisterDescriminator(Animal{}, func(a Animal) string { return a.Kind })

	m.RegisterType("dog", Dog{})
	m.RegisterType("cat", Cat{})
	m.RegisterHandlerFn(func(name string, val any) {
		fmt.Printf("%T", val)
		switch real := val.(type) {
		case Dog:
			dog = real
		case Cat:
			cat = real
		default:
			err = fmt.Errorf("didn't catch %s", name)
		}
	})

	err = m.Unmarshal([]byte(dogJson))
	assert.Nil(t, err)
	assert.Equal(t, "load", dog.Bark)

	err = m.Unmarshal([]byte(catJson))
	assert.Nil(t, err)
	assert.Equal(t, "low", cat.Meow)
}

// func Test_InterfaceMarshaler_Marshal(t *testing.T) {
// 	m := marshaler.NewWrapped()
// 	m.Register(wrappedThingA{})
// 	bs, err := m.Marshal(wrappedThingA{ThingA: "somethingA"})

// 	assert.Nil(t, err)

// 	actual := string(bs)
// 	assert.JSONEq(t, interfaceMessage, actual)
// }
