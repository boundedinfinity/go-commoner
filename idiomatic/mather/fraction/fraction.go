package fraction

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

func New[T ~int](numerator, denominator T) Fraction[T] {
	return Fraction[T]{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

// FromFloat creates a Fraction from a floating point number.
func FromFloat[T ~int, F constraints.Float](n F) Fraction[T] {
	nStr := fmt.Sprintf("%f", n)
	nStr = stringer.TrimRight(nStr, "0")
	splitOnDecimal := stringer.Split(nStr, ".")
	var size int

	if len(splitOnDecimal) > 0 {
		size = len(splitOnDecimal[1])
	}

	denominator, err := strconv.Atoi("1" + strings.Repeat("0", size))

	if err != nil {
		panic(err)
	}

	var numerator int = int(n * F(denominator))

	return Fraction[T]{
		Numerator:   T(numerator),
		Denominator: T(denominator),
	}
}

// ----------------------------------------------------------------------------------------------------
// Type
// ----------------------------------------------------------------------------------------------------

type Fraction[T ~int] struct {
	Numerator   T
	Denominator T
}

func (t Fraction[T]) String() string {
	return fmt.Sprintf("%d/%d", t.Numerator, t.Denominator)
}

func (t Fraction[T]) Float() float64 {
	fn := func(n, d float64) float64 { return n / d }
	return internal.DoubleToSingle[T, float64](t.Numerator, t.Denominator, fn)
}

func (t Fraction[T]) Copy() Fraction[T] {
	return Fraction[T]{
		Numerator:   t.Numerator,
		Denominator: t.Denominator,
	}
}

func (t Fraction[T]) Reduce() Fraction[T] {
	gcd := mather.GreatestCommonFactor(t.Numerator, t.Denominator)

	return Fraction[T]{
		Numerator:   t.Numerator / gcd,
		Denominator: t.Denominator / gcd,
	}
}

func (t Fraction[T]) IsZero() bool {
	return IsZero(t)
}

func (t Fraction[T]) Enumerate(l, h T) []Fraction[T] {
	var items []Fraction[T]
	item := t.Reduce()
	l = mather.Max(l, item.Denominator)

	for i := l; i <= h; i <<= 1 {
		items = append(items, item)
		item = New[T](item.Numerator*2, item.Denominator*2)
	}

	return items
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

func IsZero[T ~int](elem Fraction[T]) bool {
	var zero Fraction[T]
	return elem == zero
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
