package rational

import (
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"golang.org/x/exp/constraints"
)

type Rational struct {
	Integer  int
	Fraction Fraction
}

func RationalFromString[T ~string](s T) (Rational, error) {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		return Rational{}, err
	}

	return RationalFromFloat(n)
}

func RationalFromFloat[T constraints.Float](n T) (Rational, error) {
	return Rational{
		Integer:  IntegerComponent(n),
		Fraction: FractionFromFloat(n),
	}, nil
}

func IntegerComponent[T constraints.Float](x T) int {
	i, _ := math.Modf[T, int](x)
	return i
}
