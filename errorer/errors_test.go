package errorer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

func Test_Errorer(t *testing.T) {
	testCases := []struct {
		name   string
		errs   []error
		format string
		a      []any
	}{
		{
			name: "case 1",
			errs: []error{
				errorer.New("root error 1"),
				errorer.New("root error 2"),
			},
			format: "something when wrong! %s",
			a:      []any{"hmm..."},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			errRootFn := errorer.Func(tc.errs...)
			actual := errRootFn(tc.format, tc.a...)

			if len(tc.errs) > 0 {
				for _, err := range tc.errs {
					if !errorer.Is(actual, err) {
						tt.Errorf("expected error to be %v, but got %v", err, actual)
					}
				}
			}
		})
	}
}
