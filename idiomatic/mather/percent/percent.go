package percent

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
	"golang.org/x/exp/constraints"
)

func FromFloat[T ~int, F constraints.Float](v F) Percent[T] {
	return Percent[T]{
		fraction: fraction.FromFloat[T](v),
	}
}

type Percent[T ~int] struct {
	fraction fraction.Fraction[T]
}

func (this Percent[T]) Percentage() float64 {
	return this.fraction.Float()
}
