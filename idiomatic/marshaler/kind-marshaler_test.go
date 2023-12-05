package marshaler_test

import (
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

func (t Dog) TypeName() string {
	return "dog"
}

type Cat struct {
	Kind string
	Meow string
}

func (t Cat) TypeName() string {
	return "cat"
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
	m := marshaler.NewKind(Animal{}, func(a Animal) string { return a.Kind })
	m.RegisterNamer(Dog{})
	m.RegisterNamer(Cat{})

	var dog Dog
	var cat Cat
	var val any
	var err error
	var ok bool

	val, err = m.Unmarshal([]byte(dogJson))
	assert.Nil(t, err)
	dog, ok = val.(Dog)
	assert.True(t, ok)
	assert.Equal(t, "load", dog.Bark)

	val, err = m.Unmarshal([]byte(catJson))
	assert.Nil(t, err)
	cat, ok = val.(Cat)
	assert.True(t, ok)
	assert.Equal(t, "low", cat.Meow)
}
