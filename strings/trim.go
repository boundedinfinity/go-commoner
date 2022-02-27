package strings

import (
	 "strings"
	 "github.com/boundedinfinity/commons/slices"
)

func TrimSpace[T ~string](s T) T {
	return TrimSpaceLeft(TrimSpaceRight(s))
}

func TrimSpaceRight[T ~string](s T) T {
	return TrimRunesRight[T](s, ' ')
}

func TrimSpaceLeft[T ~string](s T) T {
	return TrimRunesLeft[T](s, ' ')
}

func Trim[T ~string](s T, set string) T {
	rs := string2Runes(set)
	o := TrimRunesLeft[T](s, rs...)
	return TrimRunesRight[T](o, rs...)
}

func TrimRight[T ~string](s T, set string) T {
	return TrimRunesRight[T](s, string2Runes(set)...)
}

func TrimLeft[T ~string](s T, set string) T {
	return TrimRunesLeft[T](s, string2Runes(set)...)
}

func string2Runes(s string) []rune {
	var rs []rune

	for _, r := range s {
		rs =append(rs, r)
	}

	return rs
}

func TrimRunes[T ~string](s T, set ...rune) T {
	o := TrimRunesLeft[T](s, set...)
	return TrimRunesRight[T](o, set...)
}

func TrimRunesRight[T ~string](s T, set ...rune) T {
	return TrimFuncRight[T](s, containsRune(set))
}

func TrimRunesLeft[T ~string](s T, set ...rune) T {
	return TrimFuncLeft[T](s, containsRune(set))
}

func containsRune(set []rune) func(rune) bool {
	return func(r rune) bool {
		return slices.Contains(set, r)
	}
}

func TrimPrefix[T ~string](s T, prefix string) T {
	return T(strings.TrimPrefix(string(s), prefix))
}

func TrimSuffix[T ~string](s T, suffix string) T {
	return T(strings.TrimSuffix(string(s), suffix))
}

func TrimFunc[T ~string](s T, fn func(rune) bool) T {
	o := TrimFuncLeft[T](s, fn)
	return TrimFuncRight[T](o, fn)
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