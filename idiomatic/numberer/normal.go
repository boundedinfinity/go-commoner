package numberer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

func Fraction(numerator, denominator int) Number {
	return &FractionalNumber{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

var _ Number = &FractionalNumber{}

type FractionalNumber struct {
	Numerator   int
	Denominator int
}

// ToFloat implements [Number].
func (this *FractionalNumber) ToFloat() Number {
	panic("unimplemented")
}

func (this FractionalNumber) String() string {
	return fmt.Sprintf("%d/%d", this.Numerator, this.Denominator)
}

func (this FractionalNumber) Copy() FractionalNumber {
	return FractionalNumber{
		Numerator:   this.Numerator,
		Denominator: this.Denominator,
	}
}

func (this FractionalNumber) ToFraction() Number {
	return &this
}

func (this FractionalNumber) ToNormal() Number {
	return &this
}

func (this FractionalNumber) ToMixed() Number {
	if this.IsProper() {
		return &MixedNumber{Fraction: this.Copy()}
	} else {
		return &MixedNumber{
			Whole: this.Numerator / this.Denominator,
			Fraction: FractionalNumber{
				Numerator:   this.Numerator % this.Denominator,
				Denominator: this.Denominator,
			},
		}
	}
}

func (this FractionalNumber) Reduce() FractionalNumber {
	gcd := GreatestCommonFactor(this.Numerator, this.Denominator)

	return FractionalNumber{
		Numerator:   this.Numerator / gcd,
		Denominator: this.Denominator / gcd,
	}
}

func (this FractionalNumber) IsZero() bool {
	var zero FractionalNumber
	return this == zero
}

func (this *FractionalNumber) IsImproper() bool {
	return this.Numerator > this.Denominator
}

func (this *FractionalNumber) IsProper() bool {
	return this.Denominator > this.Numerator
}

func (this *FractionalNumber) Reciprocal() Number {
	return &FractionalNumber{
		Numerator:   this.Denominator,
		Denominator: this.Numerator,
	}
}

func (t FractionalNumber) Enumerate(l, h int) []FractionalNumber {
	var items []FractionalNumber
	item := t.Reduce()
	l = mather.Max(l, item.Denominator)

	for i := l; i <= h; i <<= 1 {
		items = append(items, item)
		item = *Fraction(item.Numerator*2, item.Denominator*2).(*FractionalNumber)
	}

	return items
}

func (this FractionalNumber) Float() float64 {
	return float64(this.Numerator) / float64(this.Denominator)
}
