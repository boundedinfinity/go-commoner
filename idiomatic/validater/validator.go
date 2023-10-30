package validator

import "github.com/boundedinfinity/go-commoner/functional/mapper"

type Validator struct {
	groupMap mapper.Mapper[string, Validatable]
}

func NewValidator() *Validator {
	return &Validator{
		groupMap: make(mapper.Mapper[string, Validatable]),
	}
}

func (t Validator) Validate(groups ...string) error {
	// var validators []Validatable

	if len(groups) == 0 {
		// validators = append(validators, t.groupMap.Values()...)
	}

	return nil
}
