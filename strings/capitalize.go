package strings

import "strings"

func Uppercase[T ~string](s T) T {
	return T(strings.ToUpper(string(s)))
}

func ToUpper[T ~string](s T) T {
	return Uppercase(s)
}

func Lowercase[T ~string](s T) T {
	return T(strings.ToLower(string(s)))
}

func ToLower[T ~string](s T) T {
	return  Lowercase(s)
}

func Title[T ~string](s T) T {
	return T(strings.Title(string(s)))
}

func Capitalize[T ~string](s T) T {
	return Title(s)
}