package environmenter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/environmenter"
	"github.com/stretchr/testify/assert"
)

func Test_Environmenter(t *testing.T) {
	tcs := []struct {
		name        string
		environment []string
		inputs      []string
		expected    []string
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
			actual := environmenter.New().WithEnvironment(tc.environment...).SubAll(tc.inputs...)
			assert.EqualValues(t, actual, tc.expected)
		})
	}
}
