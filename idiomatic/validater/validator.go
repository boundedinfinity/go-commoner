package validater

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type Validation interface {
	Groups() []string
	InGroup(string) bool
	Validate(any) error
}

type validator[T any] struct {
	groups []string
}

func (t *validator[T]) Groups() []string {
	return t.groups
}

func (t *validator[T]) InGroup(name string) bool {
	return slicer.Contains(name, t.groups...)
}

func (t *validator[T]) AddGroups(groups ...string) {
	t.groups = append(t.groups, groups...)
}

//---------------------------------------------------------------------------------

var strings = stringValidator{}

type stringValidator struct{}

// func (v stringValidator) Required() Validation {

// }

//---------------------------------------------------------------------------------

func New[T any]() *Validater[T] {
	return &Validater[T]{}
}

type Validater[T any] struct {
	validations []Validation
}

func (t Validater[T]) Validate(value T, groups ...string) error {
	var validations []Validation

	if len(groups) == 0 {
		validations = t.validations
	} else {
		validations = slicer.Filter(func(_ int, validation Validation) bool {
			return slicer.HasIntersection(groups, validation.Groups())
		}, t.validations...)
	}

	for _, validation := range validations {
		if err := validation.Validate(value); err != nil {
			return err
		}
	}

	return nil
}
