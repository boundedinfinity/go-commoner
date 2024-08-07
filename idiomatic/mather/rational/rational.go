package rational

import (
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/fraction"
	"golang.org/x/exp/constraints"
)

var (
	zero_rational = Rational{}
)

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

func New(whole, numerator, denominator int) Rational {
	return Rational{
		Whole:    whole,
		Fraction: fraction.New(numerator, denominator),
	}
}

func FromString[T ~string](s T) (Rational, error) {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		return Rational{}, err
	}

	return FromFloat(n), nil
}

// ----------------------------------------------------------------------------------------------------
// Type
// ----------------------------------------------------------------------------------------------------

type Rational struct {
	Whole    int
	Fraction fraction.Fraction
}

func (t Rational) String() string {
	var result string

	if t.Whole != 0 {
		result = strconv.Itoa(t.Whole)
	}

	if !t.Fraction.IsZero() {
		if result != "" {
			result += " "
		}

		result += t.Fraction.String()
	}

	return result
}

func (t Rational) IsZero() bool {
	return IsZero(t)
}

func (t Rational) Copy() Rational {
	return Rational{
		Whole:    t.Whole,
		Fraction: t.Fraction.Copy(),
	}
}

func (t Rational) Float() float64 {
	return float64(t.Whole) + t.Fraction.Float()
}

func (t Rational) Reduce() Rational {
	return Rational{
		Whole:    t.Whole,
		Fraction: t.Fraction.Reduce(),
	}
}

func (t Rational) Mixed() Rational {
	if t.Fraction.Numerator < t.Fraction.Denominator {
		return t.Copy()
	}

	whole := t.Fraction.Numerator / t.Fraction.Denominator
	numerator := t.Fraction.Numerator % t.Fraction.Denominator

	return Rational{
		Whole: t.Whole + whole,
		Fraction: fraction.Fraction{
			Numerator:   numerator,
			Denominator: t.Fraction.Denominator,
		},
	}
}

func (t Rational) Improper() Rational {
	if t.Whole == 0 {
		return t.Copy()
	}

	return Rational{
		Whole: 0,
		Fraction: fraction.Fraction{
			Numerator:   t.Fraction.Denominator*t.Whole + t.Fraction.Numerator,
			Denominator: t.Fraction.Denominator,
		},
	}
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

func IsZero(elem Rational) bool {
	return elem == zero_rational
}

func MustString[T ~string](s T) Rational {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		panic(err)
	}

	return FromFloat(n)
}

func FromFloat[T constraints.Float](n T) Rational {
	return Rational{
		Whole:    Whole(n),
		Fraction: fraction.FromFloat(n),
	}
}

func Whole[T constraints.Float](x T) int {
	return int(x)
}

func FractionComponent[T constraints.Float](x T) int {
	return fraction.Component(x)
}

func Fraction[T constraints.Float](x T) fraction.Fraction {
	return fraction.FromFloat(x)
}
