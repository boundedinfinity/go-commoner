package slicer_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	type Thing1 struct {
		aItem string
	}

	type Thing2 struct {
		bItem string
	}

	testCases := []struct {
		name     string
		input    []Thing1
		expected []Thing2
	}{
		{
			name:     "map 1",
			input:    []Thing1{{aItem: "a"}, {aItem: "b"}},
			expected: []Thing2{{bItem: "a"}, {bItem: "b"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			fn := func(_ int, t1 Thing1) Thing2 {
				return Thing2{bItem: t1.aItem}
			}

			actual := slicer.Map(fn, tc.input...)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func Test_MapErr_NoErr(t *testing.T) {
	type Thing1 struct {
		aItem string
	}

	type Thing2 struct {
		bItem string
	}

	testCases := []struct {
		name     string
		input    []Thing1
		expected []Thing2
		err      error
		fn       func(_ int, t1 Thing1) (Thing2, error)
	}{
		{
			name:     "no error",
			input:    []Thing1{{aItem: "a"}, {aItem: "b"}},
			expected: []Thing2{{bItem: "a"}, {bItem: "b"}},
			fn: func(_ int, t1 Thing1) (Thing2, error) {
				return Thing2{bItem: t1.aItem}, nil
			},
		},
		{
			name:     "with error",
			input:    []Thing1{{aItem: "a"}, {aItem: "b"}},
			expected: []Thing2{},
			fn:       func(_ int, t1 Thing1) (Thing2, error) { return Thing2{}, errors.New("map error") },
			err:      errors.New("map error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := slicer.MapErr(tc.fn, tc.input...)
			assert.Equal(tt, tc.expected, actual)
			assert.Equal(tt, tc.err, err)
		})
	}
}
