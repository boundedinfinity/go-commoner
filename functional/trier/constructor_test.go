package trier_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/functional/trier"
	"github.com/stretchr/testify/assert"
)

func Test_trier_constructors(t *testing.T) {
	tcs := []struct {
		name              string
		actual            trier.Try[int]
		expectedResult    int
		expectedSucceeded bool
		expectedErr       error
	}{
		{
			name:              "Complete constructor, non-nil error",
			actual:            trier.Complete(1, assert.AnError),
			expectedResult:    1,
			expectedErr:       assert.AnError,
			expectedSucceeded: false,
		},
		{
			name:              "Complete constructor, nil error",
			actual:            trier.Complete(1, nil),
			expectedResult:    1,
			expectedErr:       nil,
			expectedSucceeded: true,
		},
		{
			name:              "Success constructor",
			actual:            trier.Success(1),
			expectedResult:    1,
			expectedErr:       nil,
			expectedSucceeded: true,
		},
		{
			name:              "Failure constructor",
			actual:            trier.Failure[int](assert.AnError),
			expectedResult:    0,
			expectedErr:       assert.AnError,
			expectedSucceeded: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.actual.Value, tc.name)
			assert.Equal(t, tc.expectedSucceeded, tc.actual.Succeeded(), tc.name)
			assert.Equal(t, !tc.expectedSucceeded, tc.actual.Failed(), tc.name)

			if tc.expectedErr == nil {
				assert.Nil(t, tc.actual.Err, tc.name)
			} else {
				assert.ErrorIs(t, tc.actual.Err, tc.expectedErr, tc.name)
			}
		})
	}
}

func Test_trier_lift_constructors(t *testing.T) {
	tcs := []struct {
		name              string
		actual            trier.Try[optioner.Option[int]]
		expectedResult    optioner.Option[int]
		expectedSucceeded bool
		expectedErr       error
	}{
		{
			name:              "CompleteOfZero constructor, with value, nil error",
			actual:            trier.CompleteOfZero(1, nil),
			expectedResult:    optioner.Some(1),
			expectedErr:       nil,
			expectedSucceeded: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(t, tc.expectedSucceeded, tc.actual.Succeeded(), tc.name)
			assert.Equal(t, !tc.expectedSucceeded, tc.actual.Failed(), tc.name)
			assert.Equal(t, tc.expectedResult, tc.actual.Value, tc.name)

			if tc.expectedErr == nil {
				assert.Nil(t, tc.actual.Err, tc.name)
			} else {
				assert.ErrorIs(t, tc.actual.Err, tc.expectedErr, tc.name)
			}
		})
	}
}
