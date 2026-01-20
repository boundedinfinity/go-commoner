package numberer

type Numberer interface {
	ToFraction() Fraction
}

type Fraction interface {
	ToMixed() MixedFraction
	ToCommon() CommonFraction
	IsProper() bool
	IsImproper() bool
	Reciprocal() Fraction
	Float() float64
}
