package environmenter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/environmenter"
	"github.com/stretchr/testify/assert"
)

func Test_Environmenter_WithLines(t *testing.T) {
	tcs := []struct {
		name        string
		environment []string
		inputs      []string
		expected    []string
		err         error
	}{
		{
			name:        "sub with defaults",
			environment: []string{"THING=/thing"},
			inputs: []string{
				"$THING/dir1",
				"${THING}/dir2",
				"${env:THING}/dir3",
			},
			expected: []string{
				"/thing/dir1",
				"/thing/dir2",
				"/thing/dir3",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			env := environmenter.New().WithLines("testing", tc.environment...)
			assert.Equal(tt, tc.err, env.Process())
			actual := env.SubstitueAll(tc.inputs...)
			assert.EqualValues(t, actual, tc.expected)
		})
	}
}

func Test_Environmenter_WithData(t *testing.T) {
	tcs := []struct {
		name        string
		environment string
		inputs      []string
		expected    []string
		err         error
	}{
		{
			name:        "sub with defaults",
			environment: `THING=/thing`,
			inputs: []string{
				"$THING/dir1",
				"${THING}/dir2",
				"${env:THING}/dir3",
			},
			expected: []string{
				"/thing/dir1",
				"/thing/dir2",
				"/thing/dir3",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			env := environmenter.New().WithData("testing", tc.environment)
			assert.Equal(tt, tc.err, env.Process())
			actual := env.SubstitueAll(tc.inputs...)
			assert.EqualValues(t, actual, tc.expected)
		})
	}
}
