package numberer

// type Float constraints.Float
// type Integer constraints.Integer
// type Ordered constraints.Ordered

// type Number interface {
// 	Integer | Float | Fraction
// }

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
