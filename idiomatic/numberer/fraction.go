package numberer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Common Fraction
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewCommonFraction(numerator, denominator int) CommonFraction {
	return CommonFraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

var _ Numberer = &CommonFraction{}
var _ Fraction = &CommonFraction{}

type CommonFraction struct {
	Numerator   int
	Denominator int
}

func (this CommonFraction) String() string {
	return fmt.Sprintf("%d/%d", this.Numerator, this.Denominator)
}

func (this CommonFraction) Copy() CommonFraction {
	return CommonFraction{
		Numerator:   this.Numerator,
		Denominator: this.Denominator,
	}
}

func (this CommonFraction) ToFraction() Fraction {
	fraction := this.Copy()
	return &fraction
}

func (this CommonFraction) ToMixed() MixedFraction {
	if this.IsProper() {
		return MixedFraction{
			Fraction: this.Copy(),
		}
	} else {
		return MixedFraction{
			Whole: this.Numerator / this.Denominator,
			Fraction: CommonFraction{
				Numerator:   this.Numerator % this.Denominator,
				Denominator: this.Denominator,
			},
		}
	}
}

func (this CommonFraction) ToCommon() CommonFraction {
	return this
}

func (this CommonFraction) Reduce() CommonFraction {
	gcd := mather.GreatestCommonFactor(this.Numerator, this.Denominator)

	return CommonFraction{
		Numerator:   this.Numerator / gcd,
		Denominator: this.Denominator / gcd,
	}
}

func (this CommonFraction) IsZero() bool {
	var zero CommonFraction
	return this == zero
}

func (this *CommonFraction) IsImproper() bool {
	return this.Numerator > this.Denominator
}

func (this *CommonFraction) IsProper() bool {
	return this.Denominator > this.Numerator
}

func (this *CommonFraction) Reciprocal() Fraction {
	return &CommonFraction{
		Numerator:   this.Denominator,
		Denominator: this.Numerator,
	}
}

func (t CommonFraction) Enumerate(l, h int) []CommonFraction {
	var items []CommonFraction
	item := t.Reduce()
	l = mather.Max(l, item.Denominator)

	for i := l; i <= h; i <<= 1 {
		items = append(items, item)
		item = NewCommonFraction(item.Numerator*2, item.Denominator*2)
	}

	return items
}

func (this CommonFraction) Float() float64 {
	return float64(this.Numerator) / float64(this.Denominator)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Mixed Fraction
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Numberer = &MixedFraction{}

type MixedFraction struct {
	Whole    int
	Fraction CommonFraction
}

func (this MixedFraction) String() string {
	if this.Whole == 0 {
		return this.Fraction.String()
	}

	return fmt.Sprintf("%d %v", this.Whole, this.Fraction)
}

func (this MixedFraction) Copy() MixedFraction {
	return MixedFraction{
		Whole:    this.Whole,
		Fraction: this.Fraction.Copy(),
	}
}

func (this MixedFraction) ToFraction() Fraction {
	fraction := this.Copy()
	return &fraction
}

func (this MixedFraction) ToCommon() CommonFraction {
	return CommonFraction{
		Numerator:   this.Whole*this.Fraction.Denominator + this.Fraction.Numerator,
		Denominator: this.Fraction.Denominator,
	}
}

func (this MixedFraction) ToMixed() MixedFraction {
	return this.Copy()
}

func (this MixedFraction) IsImproper() bool {
	if this.Whole == 0 {
		return this.Fraction.IsImproper()
	}
	return true
}

func (this MixedFraction) IsProper() bool {
	if this.Whole == 0 {
		return this.Fraction.IsProper()
	}
	return false
}

func (this MixedFraction) Reduce() MixedFraction {
	return MixedFraction{
		Whole:    this.Whole,
		Fraction: this.Fraction.Reduce(),
	}
}

func (t MixedFraction) Enumerate(l, h int) []MixedFraction {
	var items []MixedFraction
	fractions := t.Fraction.Enumerate(l, h)

	for _, fraction := range fractions {
		items = append(items, MixedFraction{
			Whole:    t.Whole,
			Fraction: fraction,
		})
	}

	return items
}

func (this MixedFraction) Reciprocal() Fraction {
	fraction := this.ToCommon()
	return fraction.Reciprocal()
}

func (this MixedFraction) Float() float64 {
	fraction := this.ToCommon()
	return float64(fraction.Numerator) / float64(fraction.Denominator)
}
