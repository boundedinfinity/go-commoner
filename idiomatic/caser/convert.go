package caser

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func ConverterCombinations() []string {
	outputs := []string{}

	for _, c1 := range Cases.Values() {
		for _, c2 := range Cases.Values() {
			if c1 == c2 {
				continue
			}

			if c1 == Cases.Unknown || c2 == Cases.Unknown {
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
	from, err := Cases.Parse(fromStr)

	if err != nil {
		return nil, err
	}

	toStr := stringer.TrimSpace(comps[1])
	to, err := Cases.Parse(toStr)

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

func Convert[V ~string](v V, from, to Case) string {
	var splitFn func(string) []string
	var mapFn *slicer.Pipe[string]
	var joinFn func([]string) string
	var o string

	switch from {
	case Cases.Camel:
		splitFn = splitOnCapitalOrNumber
	case Cases.Pascal:
		splitFn = splitOnCapitalOrNumber
	case Cases.Snake, Cases.SnakeLower, Cases.SnakeUpper:
		splitFn = splitOnUnderscore
	case Cases.Kebab, Cases.KebabLower, Cases.KebabUpper:
		splitFn = splitOnDash
	case Cases.Phrase:
		splitFn = splitOnSpace
	default:
		o = string(v)
	}

	wrap := func(fn func(string) string) func(int, string) string {
		return func(_ int, s string) string {
			return fn(s)
		}
	}

	switch to {
	case Cases.Camel:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToLower[string])).Then(wrap(stringer.ToCapital[string]))
		joinFn = joinWithNoSpace
	case Cases.Phrase:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToLower[string])).Then(wrap(stringer.ToCapital[string]))
		joinFn = joinWithSpace
	case Cases.Snake, Cases.SnakeLower:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToLower[string]))
		joinFn = joinWithUnderscore
	case Cases.SnakeUpper:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToUpper[string]))
		joinFn = joinWithUnderscore
	case Cases.Pascal:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToLower[string])).Then(wrap(stringer.ToCapital[string]))
		joinFn = joinWithNoSpace
	case Cases.Kebab, Cases.KebabLower:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToLower[string]))
		joinFn = joinWithDash
	case Cases.KebabUpper:
		mapFn = slicer.NewPipe[string]().Then(wrap(stringer.ToUpper[string]))
		joinFn = joinWithDash
	default:
		o = string(v)
	}

	if splitFn != nil && mapFn != nil && joinFn != nil {
		o = splitMapJoin(string(v), splitFn, mapFn, joinFn)

		switch to {
		case Cases.Camel:
			o = stringer.ToLowerFirst(o)
		}
	}

	return o
}

// Camel to ...

func CamelToPhrase[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.Phrase)
}

func CamelToSnake[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.Snake)
}

func CamelToSnakeLower[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.SnakeLower)
}

func CamelToSnakeUppser[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.SnakeUpper)
}

func CamelToPascal[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.Pascal)
}

func CamelToKebab[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.Kebab)
}

func CamelToKebabLower[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.KebabLower)
}

func CamelToKebabUpper[V ~string](v V) string {
	return Convert(v, Cases.Camel, Cases.KebabUpper)
}

// Pascal to ...

func PascalToPhrase[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.Phrase)
}

func PascalToCamel[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.Camel)
}

func PascalToSnake[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.Snake)
}

func PascalToSnakeLower[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.SnakeLower)
}

func PascalToSnakeUpper[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.SnakeUpper)
}

func PascalToKebab[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.Kebab)
}

func PascalToKebabLower[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.KebabLower)
}

func PascalToKebabUpper[V ~string](v V) string {
	return Convert(v, Cases.Pascal, Cases.KebabUpper)
}

// Snake to ...

func SnakeToPhrase[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.Phrase)
}

func SnakeToCamel[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.Camel)
}

func SnakeToPascal[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.Pascal)
}

func SnakeToKebab[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.Kebab)
}

func SnakeToKebabLower[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.KebabLower)
}

func SnakeToKebabUpper[V ~string](v V) string {
	return Convert(v, Cases.Snake, Cases.KebabUpper)
}

// Kebab to ...

func KebabToPhrase[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.Phrase)
}

func KebabToSnake[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.Snake)
}

func KebabToSnakeLower[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.SnakeLower)
}

func KebabToSnakeUpper[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.SnakeUpper)
}

func KebabToCamel[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.Camel)
}

func KebabToPascal[V ~string](v V) string {
	return Convert(v, Cases.Kebab, Cases.Pascal)
}

// Phrase to ...

func PhraseToCamel[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.Camel)
}

func PhraseToPascal[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.Pascal)
}

func PhraseToSnake[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.Snake)
}

func PhraseToSnakeLower[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.SnakeLower)
}

func PhraseToSnakeUpper[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.SnakeUpper)
}

func PhraseToKebab[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.Kebab)
}

func PhraseToKebabLower[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.KebabLower)
}

func PhraseToKebabUpper[V ~string](v V) string {
	return Convert(v, Cases.Phrase, Cases.KebabUpper)
}
