package optioner_test

import (
	"encoding/json"
	"testing"

	o "github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_serialize(t *testing.T) {
	type something[T any] struct {
		V o.Option[T] `json:",omitempty"`
	}

	testCases := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name:     "some string",
			input:    o.Some("s"),
			expected: `"s"`,
		},
		{
			name:     "none string",
			input:    o.None[string](),
			expected: `null`,
		},
		{
			name:     "some int",
			input:    o.Some(1),
			expected: `1`,
		},
		{
			name:     "none int",
			input:    o.None[int](),
			expected: `null`,
		},
		{
			name:     "some string",
			input:    something[string]{V: o.Some("s")},
			expected: `{"V": "s"}`,
		},
		{
			name:     "some string",
			input:    something[string]{V: o.None[string]()},
			expected: `{"V": null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)

			assert.Equal(tt, tc.err, err)
			assert.JSONEqf(tt, tc.expected, string(actual), "%v == %v", tc.expected, string(actual))
		})
	}
}

func Test_deserialize_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected o.Option[string]
		err      error
	}{
		{
			name:     "some string",
			input:    `"s"`,
			expected: o.Some("s"),
		},
		{
			name:     "zero string",
			input:    `""`,
			expected: o.Some(""),
		},
		{
			name:     "null string",
			input:    `null`,
			expected: o.None[string](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual o.Option[string]
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, tc.expected.Get(), actual.Get(), "Get()")
			assert.Equal(tt, tc.expected.Empty(), actual.Empty(), "Empty()")
			assert.Equal(tt, tc.expected.Defined(), actual.Defined(), "Defined()")
		})
	}
}

func Test_deserialize_int(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected o.Option[int]
		err      error
	}{
		{
			name:     "some int",
			input:    `1`,
			expected: o.Some(1),
		},
		{
			name:     "zero int",
			input:    `0`,
			expected: o.Some(0),
		},
		{
			name:     "null int",
			input:    `null`,
			expected: o.None[int](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual o.Option[int]
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, tc.expected.Get(), actual.Get(), "Get()")
			assert.Equal(tt, tc.expected.Empty(), actual.Empty(), "Empty()")
			assert.Equal(tt, tc.expected.Defined(), actual.Defined(), "Defined()")
		})
	}
}

func Test_deserialize_bool(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected o.Option[bool]
		err      error
	}{
		{
			name:     "some bool",
			input:    `true`,
			expected: o.Some(true),
		},
		{
			name:     "zero bool",
			input:    `false`,
			expected: o.Some(false),
		},
		{
			name:     "null bool",
			input:    `null`,
			expected: o.None[bool](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual o.Option[bool]
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, tc.expected.Get(), actual.Get(), "Get()")
			assert.Equal(tt, tc.expected.Empty(), actual.Empty(), "Empty()")
			assert.Equal(tt, tc.expected.Defined(), actual.Defined(), "Defined()")
		})
	}
}

func Test_deserialize_struct(t *testing.T) {
	type something[T string] struct {
		V o.Option[T] `json:",omitempty"`
	}

	testCases := []struct {
		name     string
		input    string
		expected something[string]
		err      error
	}{
		{
			name:     "some struct",
			input:    `{"V": "s"}`,
			expected: something[string]{V: o.Some("s")},
		},
		{
			name:     "zero struct",
			input:    `{}`,
			expected: something[string]{V: o.None[string]()},
		},
		{
			name:     "null struct",
			input:    `{"V": null}`,
			expected: something[string]{V: o.None[string]()},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			var actual something[string]
			err := json.Unmarshal([]byte(tc.input), &actual)

			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, tc.expected.V.Get(), actual.V.Get(), "Get()")
			assert.Equal(tt, tc.expected.V.Empty(), actual.V.Empty(), "Empty()")
			assert.Equal(tt, tc.expected.V.Defined(), actual.V.Defined(), "Defined()")
		})
	}
}
