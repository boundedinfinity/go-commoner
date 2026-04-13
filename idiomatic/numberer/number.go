package numberer

type Number interface {
	IsZero() bool
	ToFloat() Number
	ToFraction() Number
	// ToMixed() Number
	// ToImproper() Number
	// IsProper() bool
	// IsImproper() bool
	// Reciprocal() Number
	// Float() float64
}
