package caser

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Convert[V ~string](v V, from, to CaseType) string {
	var splitFn func(string) []string
	var mapFn *slicer.Pipe[string]
	var joinFn func([]string) string
	var o string

	switch from {
	case CaseTypes.Camel:
		splitFn = splitOnCapitalOrNumber
	case CaseTypes.Pascal:
		splitFn = splitOnCapitalOrNumber
	case CaseTypes.Snake, CaseTypes.Snakeupper:
		splitFn = splitOnUnderscore
	case CaseTypes.Kebab, CaseTypes.Kebabupper:
		splitFn = splitOnDash
	case CaseTypes.Phrase:
		splitFn = splitOnSpace
	default:
		o = string(v)
	}

	switch to {
	case CaseTypes.Camel:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string]).Then(stringer.Title[string])
		joinFn = joinWithNoSpace
	case CaseTypes.Phrase:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string]).Then(stringer.Title[string])
		joinFn = joinWithSpace
	case CaseTypes.Snake:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string])
		joinFn = joinWithUnderscore
	case CaseTypes.Snakeupper:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToUpper[string])
		joinFn = joinWithUnderscore
	case CaseTypes.Pascal:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string]).Then(stringer.Title[string])
		joinFn = joinWithNoSpace
	case CaseTypes.Kebab:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string])
		joinFn = joinWithDash
	case CaseTypes.Kebabupper:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToUpper[string])
		joinFn = joinWithDash
	default:
		o = string(v)
	}

	if splitFn != nil && mapFn != nil && joinFn != nil {
		o = splitMapJoin(string(v), splitFn, mapFn, joinFn)

		switch to {
		case CaseTypes.Camel:
			o = stringer.LowercaseFirst(o)
		}
	}

	return o
}

func CamelToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Phrase)
}

func CamelToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Snake)
}

func CamelToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Pascal)
}

func CamelToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Kebab)
}

func CamelToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Kebabupper)
}

func PascalToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Phrase)
}

func PascalToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Camel)
}

func PascalToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Snake)
}

func PascalToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Snakeupper)
}

func PascalToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Kebab)
}

func PascalToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Kebabupper)
}

func SnakeToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.Phrase)
}

func SnakeToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.Camel)
}

func SnakeToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.Pascal)
}

func SnakeToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.Kebab)
}

func SnakeToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.Kebabupper)
}

func KebabToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Phrase)
}

func KebabToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Snake)
}

func KebabToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Snakeupper)
}

func KebabToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Camel)
}

func KebabToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Pascal)
}

func PhraseToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Camel)
}

func PhraseToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Pascal)
}

func PhraseToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Snake)
}

func PhraseToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Snakeupper)
}

func PhraseToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Kebab)
}

func PhraseToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Kebabupper)
}
