package geometry

import (
	gomath "math"

	"github.com/boundedinfinity/go-commoner/idiomatic/math/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/types"
)

type Angle[T types.Numbers] struct {
	Magnitude T
	Type      AngleType
	Direction AngleDirection
}

func (t Angle[T]) ToRadian() Angle[T] {
	switch t.Type {
	case AngleTypes.Radians:
		return t
	default:
		return Angle[T]{
			Type:      AngleTypes.Radians,
			Direction: t.Direction,
			Magnitude: DegreeToRadian(t.Magnitude),
		}
	}
}

func (t Angle[T]) ToDegree() Angle[T] {
	switch t.Type {
	case AngleTypes.Degrees:
		return t
	default:
		return Angle[T]{
			Type:      AngleTypes.Degrees,
			Direction: t.Direction,
			Magnitude: RadianToDegree(t.Magnitude),
		}
	}
}

func DegreeToRadian[T types.Numbers](angle T) T {
	return internal.SingleToSingle[T, T](angle, func(n float64) float64 {
		return n * gomath.Pi / 180
	})
}

func RadianToDegree[T types.Numbers](angle T) T {
	return internal.SingleToSingle[T, T](angle, func(n float64) float64 {
		return n * 180 / gomath.Pi
	})
}
