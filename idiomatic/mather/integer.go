package mather

import (
	"fmt"
	"strconv"
)

func OrdinalIndicator[I Integer](n I) string {
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

func Ordinalize[I Integer](n I) string {
	// https://en.wikipedia.org/wiki/Ordinal_numeral

	return fmt.Sprintf("%d%s", n, OrdinalIndicator(n))
}

func Itoa[I Integer](i I) string {
	return strconv.FormatInt(int64(i), 10)
}
