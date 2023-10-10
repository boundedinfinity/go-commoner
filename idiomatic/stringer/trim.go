package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func TrimSpace[T ~string](s T) T {
	return TrimSpaceLeft(TrimSpaceRight(s))
}

func TrimSpaceRight[T ~string](s T) T {
	return TrimRunesRight(s, ' ')
}

func TrimSpaceLeft[T ~string](s T) T {
	return TrimRunesLeft(s, ' ')
}

func Trim[T ~string](s T, set string) T {
	runes := string2Runes(set)
	res := s
	res = TrimRunesLeft(res, runes...)
	res = TrimRunesRight(res, runes...)
	return res
}

func TrimRight[T ~string](s T, set string) T {
	return TrimRunesRight(s, string2Runes(set)...)
}

func TrimLeft[T ~string](s T, set string) T {
	return TrimRunesLeft(s, string2Runes(set)...)
}

func string2Runes(s string) []rune {
	var rs []rune

	for _, r := range s {
		rs = append(rs, r)
	}

	return rs
}

func TrimRunes[T ~string](s T, set ...rune) T {
	o := TrimRunesLeft(s, set...)
	return TrimRunesRight(o, set...)
}

func TrimRunesRight[T ~string](s T, set ...rune) T {
	return TrimFuncRight(s, containsRune(set))
}

func TrimRunesLeft[T ~string](s T, set ...rune) T {
	return TrimFuncLeft(s, containsRune(set))
}

func containsRune(set []rune) func(rune) bool {
	return func(r rune) bool {
		return slicer.Contains(set, r)
	}
}

func TrimPrefix[T ~string](s T, prefix string) T {
	return T(strings.TrimPrefix(string(s), prefix))
}

func TrimSuffix[T ~string](s T, suffix string) T {
	return T(strings.TrimSuffix(string(s), suffix))
}

func TrimFunc[T ~string](s T, fn func(rune) bool) T {
	o := TrimFuncLeft(s, fn)
	return TrimFuncRight(o, fn)
}

func TrimFuncLeft[T ~string](s T, fn func(rune) bool) T {
	start := 0

	for i, x := range s {
		if !fn(rune(x)) {
			start = i
			break
		}
	}

	return s[start:]
}

func TrimFuncRight[T ~string](s T, fn func(rune) bool) T {
	end := len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		x := s[i]

		if !fn(rune(x)) {
			end = i + 1
			break
		}
	}

	return s[:end]
}
