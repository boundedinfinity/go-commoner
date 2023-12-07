package reflecter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/stretchr/testify/assert"
)

type aString string

var (
	an_input = aString("an input")
)

func Test_Instances_QualifiedName(t *testing.T) {
	actual := reflecter.Instances.QualifiedName(an_input)
	assert.Equal(t, "github.com/boundedinfinity/go-commoner/idiomatic/reflecter_test/aString", actual)
}

func Test_Instances_SimpleName(t *testing.T) {
	actual := reflecter.Instances.SimpleName(an_input)
	assert.Equal(t, "aString", actual)
}

// func Test_Errorer_constructors(t *testing.T) {
// 	testError := errorer.Errorf(assert.AnError.Error())

// 	tcs := []struct {
// 		name   string
// 		actual error
// 		is     error
// 		str    string
// 		err    string
// 	}{
// 		{
// 			name:   "New constructor",
// 			actual: errorer.Wrap(assert.AnError),
// 			is:     assert.AnError,
// 			str:    assert.AnError.Error(),
// 			err:    assert.AnError.Error(),
// 		},
// 		{
// 			name:   "Errorf constructor",
// 			actual: testError,
// 			is:     testError,
// 			str:    assert.AnError.Error(),
// 			err:    assert.AnError.Error(),
// 		},
// 		// {
// 		// 	name:   "None constructor",
// 		// 	actual: errorer.None(),
// 		// 	is:     &errorer.Errorer{},
// 		// 	str:    "",
// 		// 	err:    "",
// 		// },
// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.name, func(tt *testing.T) {
// 			assert.Error(t, tc.actual, tc.name)
// 			assert.ErrorIs(t, tc.actual, tc.is, tc.name)
// 			assert.Equal(t, tc.str, tc.actual.(*errorer.Errorer).String(), tc.name)
// 			assert.Equal(t, tc.err, tc.actual.Error(), tc.name)
// 		})
// 	}
// }
