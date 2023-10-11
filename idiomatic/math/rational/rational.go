package rational

import "strconv"

type Fraction struct {
	Numerator   int
	Demoninator int
}

type Rational struct {
	Integer  int
	Fraction *Fraction
}

func ParseRational[T ~string](s T) (Rational, error) {
	n, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		return Rational{}, err
	}

	return Rational{
		Integer: IntegerComponent(n),
	}, nil
}
