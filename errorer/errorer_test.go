package errorer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/stretchr/testify/assert"
)

func Test_Errorer_constructors(t *testing.T) {
	testError := errorer.Errorf(assert.AnError.Error())

	tcs := []struct {
		name   string
		actual error
		is     error
		str    string
		err    string
	}{
		{
			name:   "New constructor",
			actual: errorer.New(assert.AnError),
			is:     assert.AnError,
			str:    assert.AnError.Error(),
			err:    assert.AnError.Error(),
		},
		{
			name:   "Errorf constructor",
			actual: testError,
			is:     testError,
			str:    assert.AnError.Error(),
			err:    assert.AnError.Error(),
		},
		// {
		// 	name:   "None constructor",
		// 	actual: errorer.None(),
		// 	is:     &errorer.Errorer{},
		// 	str:    "",
		// 	err:    "",
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Error(t, tc.actual, tc.name)
			assert.ErrorIs(t, tc.actual, tc.is, tc.name)
			assert.Equal(t, tc.str, tc.actual.(*errorer.Errorer).String(), tc.name)
			assert.Equal(t, tc.err, tc.actual.Error(), tc.name)
		})
	}
}
