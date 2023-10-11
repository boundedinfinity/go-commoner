package rational

import (
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

type Fraction struct {
	Numerator   int
	Demoninator int
}

func (t Fraction) Reduce() Fraction {
	gcd := math.GreatestCommonDivisor(t.Numerator, t.Demoninator)

	return Fraction{
		Numerator:   t.Numerator / gcd,
		Demoninator: t.Demoninator / gcd,
	}
}

func FractionFromFloat[T constraints.Float](n T) Fraction {
	numerator := FractionComponent(n)
	size := FractionSize(n)
	denominator := int(math.Pow10(size))

	return Fraction{
		Numerator:   numerator,
		Demoninator: denominator,
	}
}

func FractionComponent[T constraints.Float](n T) int {
	s := fmt.Sprintf("%v", n)
	comps := stringer.Split(s, ".")

	if len(comps) > 1 {
		s = stringer.Split(s, ".")[1]
	} else {
		s = "0"
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func FractionSize[T constraints.Float](n T) int {
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
