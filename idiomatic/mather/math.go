package mather

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

// CountDigitsRightOfDecimal returns the number of digits after the decimal in a floating point number
//
// This will ignore non-significat trailing zeros
func CountDigitsRightOfDecimal[F types.Float](f F) int {
	var size int
	s := fmt.Sprintf("%f", f)
	c := strings.Split(s, ".")

	switch len(c) {
	case 1:
		// no decimal
	case 2:
		after := strings.TrimRight(c[1], "0")
		size = len(after)
	default:
		panic("can't process: " + s)
	}

	return size
}

func PlaceOfDigitsAfterDecimal[F types.Float](f F) int {
	size := CountDigitsRightOfDecimal(f)
	place, err := strconv.Atoi("1" + strings.Repeat("0", size))

	if err != nil {
		panic(err)
	}

	return place
}

// Signbit reports whether n is negative or negative zero.
func Signbit[N types.Number](n N) bool {
	return math.Signbit(float64(n))
}

// IsNegative reports whether n is negative or negative zero.
func IsNegative[I types.Number](n I) bool {
	return Signbit(n)
}

func Itoa[I types.Integer](i I) string {
	return strconv.FormatInt(int64(i), 10)
}

func OrdinalIndicator[I types.Integer](n I) string {
	// https://en.wikipedia.org/wiki/Ordinal_indicator

	var indicator string
	s := Itoa(n)
	c := s[len(s)-1:]

	switch c {
	case "0":
		switch len(s) {
		case 2:
			indicator = "ieth"
		default:
			indicator = "th"
		}
	case "1":
		indicator = "st"
	case "2":
		indicator = "nd"
	case "3":
		indicator = "rd"
	default:
		indicator = "th"
	}

	return indicator
}

func Ordinalize[I types.Integer](n I) string {
	// https://en.wikipedia.org/wiki/Ordinal_numeral

	return fmt.Sprintf("%d%s", n, OrdinalIndicator(n))
}
