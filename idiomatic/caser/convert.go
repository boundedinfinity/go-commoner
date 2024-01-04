package caser

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func ConverterCombinations() []string {
	outputs := []string{}

	for _, c1 := range CaseTypes.Values() {
		for _, c2 := range CaseTypes.Values() {
			if c1 == c2 {
				continue
			}

			if c1 == CaseTypes.Unknown || c2 == CaseTypes.Unknown {
				continue
			}

			outputs = append(outputs, fmt.Sprintf("%v-to-%v", c1, c2))
		}
	}

	return outputs
}

func MustConverter[V ~string](v V, convert string) func(V) string {
	converter, err := Converter[V](convert)

	if err != nil {
		panic(err)
	}

	return converter
}

func Converter[V ~string](convert string) (func(V) string, error) {
	working := convert
	comps := stringer.Split(working, "-to-")

	if len(comps) != 2 {
		return nil, fmt.Errorf("invalid convert string %v", convert)
	}

	fromStr := stringer.TrimSpace(comps[0])
	from, err := CaseTypes.Parse(fromStr)

	if err != nil {
		return nil, err
	}

	toStr := stringer.TrimSpace(comps[1])
	to, err := CaseTypes.Parse(toStr)

	if err != nil {
		return nil, err
	}

	fn := func(v V) string {
		return Convert(v, from, to)
	}

	return fn, nil
}

func MustConvertAs[V ~string](v V, convert string) string {
	o, err := ConvertAs(v, convert)

	if err != nil {
		panic(err)
	}

	return o
}

func ConvertAs[V ~string](v V, convert string) (string, error) {
	fn, err := Converter[V](convert)

	if err != nil {
		return "", err
	}

	return fn(v), nil
}

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
	case CaseTypes.Snake, CaseTypes.SnakeLower, CaseTypes.SnakeUpper:
		splitFn = splitOnUnderscore
	case CaseTypes.Kebab, CaseTypes.KebabLower, CaseTypes.KebabUpper:
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
	case CaseTypes.Snake, CaseTypes.SnakeLower:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string])
		joinFn = joinWithUnderscore
	case CaseTypes.SnakeUpper:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToUpper[string])
		joinFn = joinWithUnderscore
	case CaseTypes.Pascal:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string]).Then(stringer.Title[string])
		joinFn = joinWithNoSpace
	case CaseTypes.Kebab, CaseTypes.KebabLower:
		mapFn = slicer.NewPipe[string]().Then(stringer.ToLower[string])
		joinFn = joinWithDash
	case CaseTypes.KebabUpper:
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

// Camel to ...

func CamelToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Phrase)
}

func CamelToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Snake)
}

func CamelToSnakeLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.SnakeLower)
}

func CamelToSnakeUppser[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.SnakeUpper)
}

func CamelToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Pascal)
}

func CamelToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.Kebab)
}

func CamelToKebabLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.KebabLower)
}

func CamelToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Camel, CaseTypes.KebabUpper)
}

// Pascal to ...

func PascalToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Phrase)
}

func PascalToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Camel)
}

func PascalToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Snake)
}

func PascalToSnakeLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.SnakeLower)
}

func PascalToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.SnakeUpper)
}

func PascalToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.Kebab)
}

func PascalToKebabLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.KebabLower)
}

func PascalToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Pascal, CaseTypes.KebabUpper)
}

// Snake to ...

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

func SnakeToKebabLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.KebabLower)
}

func SnakeToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Snake, CaseTypes.KebabUpper)
}

// Kebab to ...

func KebabToPhrase[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Phrase)
}

func KebabToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Snake)
}

func KebabToSnakeLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.SnakeLower)
}

func KebabToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.SnakeUpper)
}

func KebabToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Camel)
}

func KebabToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Kebab, CaseTypes.Pascal)
}

// Phrase to ...

func PhraseToCamel[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Camel)
}

func PhraseToPascal[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Pascal)
}

func PhraseToSnake[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Snake)
}

func PhraseToSnakeLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.SnakeLower)
}

func PhraseToSnakeUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.SnakeUpper)
}

func PhraseToKebab[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.Kebab)
}

func PhraseToKebabLower[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.KebabLower)
}

func PhraseToKebabUpper[V ~string](v V) string {
	return Convert(v, CaseTypes.Phrase, CaseTypes.KebabUpper)
}
