package errorer_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/stretchr/testify/assert"
)

func Test_Errorer_constructor(t *testing.T) {
	errRoot := errorer.Errorf("root error")

	tcs := []struct {
		name   string
		actual error
		target error
		str    string
		list   []string
	}{
		{
			name:   "case 1",
			actual: errRoot,
			target: errRoot,
			str:    "root error",
			list:   []string{"root error"},
		},
		{
			name:   "case 2",
			actual: errRoot.Subf("case 2"),
			target: errRoot,
			str:    "case 2 : root error",
			list:   []string{"root error", "case 2"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Error(t, tc.actual)
			assert.ErrorContains(t, tc.actual, tc.str)
			assert.Equal(tt, tc.list, tc.actual.(*errorer.Errorer).List())
			// assert.ErrorIs(t, tc.actual, tc.target)
		})
	}
}

func Test_Errorer_marshal(t *testing.T) {
	tcs := []struct {
		name     string
		input    error
		expected string
	}{
		{
			name:     "case 1",
			input:    errorer.Errorf("case 1"),
			expected: `{"errors":["case 1"]}`,
		},
		{
			name:     "case 2",
			input:    errorer.Errorf("root error").Subf("case 2"),
			expected: `{"errors":["root error","case 2"]}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)

			assert.Nil(tt, err)
			assert.JSONEq(tt, tc.expected, string(actual), string(actual))
		})
	}
}

func Test_Errorer_marshal_embedded(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "case 1",
			input:    `{"errors":["case 1"]}`,
			expected: errorer.Errorf("case 1"),
		},
		{
			name:     "case 2",
			input:    `{"errors":["root error","case 2"]}`,
			expected: errorer.Errorf("root error").Subf("case 2"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			var actual errorer.Errorer
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Nil(tt, err)
			assert.Equal(tt, tc.expected, &actual)
		})
	}
}
