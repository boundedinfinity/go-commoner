package caser

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/boundedinfinity/go-commoner/utf8"
)

var (
	re = regexp.MustCompile(`[A-Z][^A-Z]*`)
)

func splitMapJoin(v string, splitFn func(v string) []string, mapFn func(string) string, joinFn func([]string) string) string {
	s := splitFn(v)
	m := slicer.Map(s, mapFn)
	j := joinFn(m)
	return j
}

func joinWithNoSpace(v []string) string {
	return stringer.Join(v, "")
}

func joinWithSpace(v []string) string {
	return stringer.Join(v, " ")
}

func joinWithUnderscore(v []string) string {
	return stringer.Join(v, "_")
}

func joinWithDash(v []string) string {
	return stringer.Join(v, "-")
}

func splitOnSpace(v string) []string {
	return stringer.Split(v, " ")
}

func splitOnUnderscore(v string) []string {
	return stringer.Split(v, "_")
}

func splitOnDash(v string) []string {
	return stringer.Split(v, "-")
}

func splitOnCapitalOrNumber(v string) []string {
	var os []string
	var t string

	for i := 0; i < len(v); i++ {
		c := v[i]

		if utf8.IsUpperCase(c) || utf8.IsNumber(c) {
			if t != "" {
				os = append(os, t)
				t = ""
			}
		}

		t = t + string(c)
	}

	if len(t) > 0 {
		os = append(os, t)
	}

	return os
}

func mapPipeline(fns ...func(string) string) func(string) string {
	return func(v string) string {
		o := v

		for _, fn := range fns {
			o = fn(o)
		}

		return o
	}
}

func mapNoOp(v string) string {
	return v
}
