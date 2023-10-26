package errorer_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/stretchr/testify/assert"
)

func Test_Errorer_marshal(t *testing.T) {
	tcs := []struct {
		name     string
		input    error
		expected string
	}{
		{
			name:     "New constructor",
			input:    errorer.New(assert.AnError),
			expected: `"assert.AnError general error for testing"`,
		},
		{
			name:     "Errorf constructor",
			input:    errorer.Errorf("some stuff here %v", 1),
			expected: `"some stuff here 1"`,
		},
		{
			name:     "Empty",
			input:    errorer.None(),
			expected: `null`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)

			assert.Nil(t, err)
			assert.JSONEq(t, tc.expected, string(actual))
		})
	}
}

func Test_Errorer_marshal_embedded(t *testing.T) {
	type thing struct {
		Val string
		Err error
	}

	tcs := []struct {
		name     string
		input    thing
		expected string
	}{
		{
			name:     "New constructor",
			input:    thing{Val: "V", Err: errorer.New(assert.AnError)},
			expected: `{"Val":"V", "Err":"assert.AnError general error for testing"}`,
		},
		{
			name:     "Errorf constructor",
			input:    thing{Val: "V", Err: errorer.Errorf("some stuff here %v", 1)},
			expected: `{"Val":"V", "Err":"some stuff here 1"}`,
		},
		{
			name:     "Empty constructor",
			input:    thing{Val: "V", Err: errorer.None()},
			expected: `{"Val":"V", "Err": null}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)

			assert.Nil(t, err)
			assert.JSONEq(t, tc.expected, string(actual))
		})
	}
}
