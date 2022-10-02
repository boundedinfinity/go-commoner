package optioner_test

import (
	"encoding/json"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_serialize(t *testing.T) {
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)
			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}

func Test_deserialize(t *testing.T) {
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
			name:     "no string",
			input:    `null`,
			expected: o.None[string](),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			input := []byte(tc.input)
			var actual o.Option[string]
			err := json.Unmarshal(input, &actual)

			assert.Equal(tt, tc.err, err)
			assert.Equal(tt, tc.expected.Get(), actual.Get())
			assert.Equal(tt, tc.expected.Empty(), actual.Empty())
			assert.Equal(tt, tc.expected.Defined(), actual.Defined())

		})
	}
}

// func Test_JSON_deserialize_string(t *testing.T) {
// 	input := []byte(`"s"`)
// 	expected := optioner.Some("s")
// 	var actual optioner.Option[string]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }

// func Test_JSON_deserialize_nil_string(t *testing.T) {
// 	input := []byte(`null`)
// 	expected := optioner.None[string]()
// 	var actual optioner.Option[string]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }

// func Test_JSON_deserialize_int(t *testing.T) {
// 	input := []byte(`null`)
// 	expected := optioner.None[string]()
// 	var actual optioner.Option[string]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }

// func Test_JSON_deserialize_int_empty(t *testing.T) {
// 	input := []byte(`null`)
// 	expected := optioner.None[string]()
// 	var actual optioner.Option[string]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }

// func Test_JSON_deserialize_nil_int(t *testing.T) {
// 	input := []byte(`null`)
// 	expected := optioner.None[int]()
// 	var actual optioner.Option[int]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }

// func Test_JSON_serialize_struct(t *testing.T) {
// 	type Struct1 struct {
// 		I int
// 		S string
// 	}

// 	s1 := Struct1{
// 		I: 1,
// 		S: "s",
// 	}

// 	input1 := optioner.Some(s1)
// 	expected1 := []byte(`{"I":1,"S":"s"}`)
// 	actual1, err := json.Marshal(input1)

// 	assert.Nil(t, err)
// 	assert.Equal(t, string(expected1), string(actual1))

// 	input2 := optioner.None[Struct1]()
// 	expected2 := []byte(`null`)
// 	actual2, err := json.Marshal(input2)

// 	assert.Nil(t, err)
// 	assert.Equal(t, string(expected2), string(actual2))

// 	type Struct2 struct {
// 		I optioner.Option[int]
// 		S optioner.Option[string]
// 	}

// 	input3 := Struct2{
// 		I: optioner.None[int](),
// 		S: optioner.None[string](),
// 	}

// 	expected3 := []byte(`{"I":null,"S":null}`)
// 	actual3, err := json.Marshal(input3)

// 	assert.Nil(t, err)
// 	assert.Equal(t, string(expected3), string(actual3))
// }

// func Test_JSON_deserialize_struct(t *testing.T) {
// 	input := []byte(`null`)
// 	expected := optioner.None[int]()
// 	var actual optioner.Option[int]
// 	err := json.Unmarshal(input, &actual)

// 	assert.Nil(t, err)
// 	assert.Equal(t, expected.Get(), actual.Get())
// 	assert.Equal(t, expected.Empty(), actual.Empty())
// 	assert.Equal(t, expected.Defined(), actual.Defined())
// }
