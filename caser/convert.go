package caser

import (
	"github.com/boundedinfinity/go-commoner/caser/case_type"
	"github.com/boundedinfinity/go-commoner/stringer"
)

func Convert[V ~string](v V, f, t case_type.CaseType) string {
	var splitFn func(string) []string
	var mapFn func(string) string
	var joinFn func([]string) string
	var o string

	switch f {
	case case_type.Camel:
		splitFn = splitOnCapitalOrNumber
	case case_type.Pascal:
		splitFn = splitOnCapitalOrNumber
	case case_type.Snake, case_type.Snakeupper:
		splitFn = splitOnUnderscore
	case case_type.Kebab, case_type.Kebabupper:
		splitFn = splitOnDash
	case case_type.Phrase:
		splitFn = splitOnSpace
	default:
		o = string(v)
	}

	switch t {
	case case_type.Camel:
		mapFn = mapPipeline(stringer.ToLower[string], stringer.Title[string])
		joinFn = joinWithNoSpace
	case case_type.Phrase:
		mapFn = mapPipeline(stringer.ToLower[string], stringer.Title[string])
		joinFn = joinWithSpace
	case case_type.Snake:
		mapFn = stringer.ToLower[string]
		joinFn = joinWithUnderscore
	case case_type.Snakeupper:
		mapFn = stringer.ToUpper[string]
		joinFn = joinWithUnderscore
	case case_type.Pascal:
		mapFn = mapPipeline(stringer.ToLower[string], stringer.Title[string])
		joinFn = joinWithNoSpace
	case case_type.Kebab:
		mapFn = stringer.ToLower[string]
		joinFn = joinWithDash
	case case_type.Kebabupper:
		mapFn = stringer.ToUpper[string]
		joinFn = joinWithDash
	default:
		o = string(v)
	}

	if splitFn != nil && mapFn != nil && joinFn != nil {
		o = splitMapJoin(string(v), splitFn, mapFn, joinFn)

		switch t {
		case case_type.Camel:
			o = stringer.LowercaseFirst(o)
		}
	}

	return o
}

func CamelToPhrase[V ~string](v V) string {
	return Convert(v, case_type.Camel, case_type.Phrase)
}

func CamelToSnake[V ~string](v V) string {
	return Convert(v, case_type.Camel, case_type.Snake)
}

func CamelToPascal[V ~string](v V) string {
	return Convert(v, case_type.Camel, case_type.Pascal)
}

func CamelToKebab[V ~string](v V) string {
	return Convert(v, case_type.Camel, case_type.Kebab)
}

func CamelToKebabUpper[V ~string](v V) string {
	return Convert(v, case_type.Camel, case_type.Kebabupper)
}

func PascalToPhrase[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Phrase)
}

func PascalToCamel[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Camel)
}

func PascalToSnake[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Snake)
}

func PascalToSnakeUpper[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Snakeupper)
}

func PascalToKebab[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Kebab)
}

func PascalToKebabUpper[V ~string](v V) string {
	return Convert(v, case_type.Pascal, case_type.Kebabupper)
}

func SnakeToPhrase[V ~string](v V) string {
	return Convert(v, case_type.Snake, case_type.Phrase)
}

func SnakeToCamel[V ~string](v V) string {
	return Convert(v, case_type.Snake, case_type.Camel)
}

func SnakeToPascal[V ~string](v V) string {
	return Convert(v, case_type.Snake, case_type.Pascal)
}

func SnakeToKebab[V ~string](v V) string {
	return Convert(v, case_type.Snake, case_type.Kebab)
}

func SnakeToKebabUpper[V ~string](v V) string {
	return Convert(v, case_type.Snake, case_type.Kebabupper)
}

func KebabToPhrase[V ~string](v V) string {
	return Convert(v, case_type.Kebab, case_type.Phrase)
}

func KebabToSnake[V ~string](v V) string {
	return Convert(v, case_type.Kebab, case_type.Snake)
}

func KebabToSnakeUpper[V ~string](v V) string {
	return Convert(v, case_type.Kebab, case_type.Snakeupper)
}

func KebabToCamel[V ~string](v V) string {
	return Convert(v, case_type.Kebab, case_type.Camel)
}

func KebabToPascal[V ~string](v V) string {
	return Convert(v, case_type.Kebab, case_type.Pascal)
}

func PhraseToCamel[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Camel)
}

func PhraseToPascal[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Pascal)
}

func PhraseToSnake[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Snake)
}

func PhraseToSnakeUpper[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Snakeupper)
}

func PhraseToKebab[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Kebab)
}

func PhraseToKebabUpper[V ~string](v V) string {
	return Convert(v, case_type.Phrase, case_type.Kebabupper)
}
