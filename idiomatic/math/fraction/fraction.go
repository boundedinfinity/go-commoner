package fraction

import (
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func (t Fraction) Float() int64 {
	return int64(t.Numerator) / int64(t.Denominator)
}

func (t Fraction) Copy() Fraction {
	return Fraction{
		Numerator:   t.Numerator,
		Denominator: t.Denominator,
	}
}

func (t Fraction) Reduce() Fraction {
	gcd := math.GreatestCommonFactor(t.Numerator, t.Denominator)

	return Fraction{
		Numerator:   t.Numerator / gcd,
		Denominator: t.Denominator / gcd,
	}
}

func New(numerator, denominator int) Fraction {
	return Fraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func FromFloat[T constraints.Float](n T) Fraction {
	numerator := Component(n)
	size := Magnitude(n)
	denominator := int(math.Pow10(size))

	return Fraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func Component[T constraints.Float](n T) int {
	s := fmt.Sprintf("%v", n)
	comps := stringer.Split(s, ".")
	var i int

	if len(comps) > 1 {
		s = comps[1]
		x, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		i = x
	} else {
		i = 0
	}

	return i
}

func Magnitude[T constraints.Float](n T) int {
	s := fmt.Sprintf("%v", n)
	comps := stringer.Split(s, ".")
	var size int

	if len(comps) > 1 {
		s = stringer.Split(s, ".")[1]
		size = len(s)
	} else {
		s = "0"
		size = 0
	}

	return size
}
