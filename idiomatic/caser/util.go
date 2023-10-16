package caser

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/go-commoner/utf8"
)

func splitMapJoin(v string, splitFn func(v string) []string, mapFn *slicer.Pipe[string], joinFn func([]string) string) string {
	s := splitFn(v)
	m := mapFn.List(s...)
	j := joinFn(m)
	return j
}

func joinWithNoSpace(v []string) string {
	return stringer.Join("", v...)
}

func joinWithSpace(v []string) string {
	return stringer.Join(" ", v...)
}

func joinWithUnderscore(v []string) string {
	return stringer.Join("_", v...)
}

func joinWithDash(v []string) string {
	return stringer.Join("-", v...)
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
