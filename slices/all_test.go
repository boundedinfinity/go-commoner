package slices_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/slices"
	"github.com/stretchr/testify/assert"
)

func Test_All_String(t *testing.T) {
	actual1 := slices.AllV("a", "a", "a", "a")
	assert.Equal(t, true, actual1)

	actual2 := slices.AllV("a", "b", "a", "a")
	assert.Equal(t, false, actual2)
}

// func Test_All_UserDefined(t *testing.T) {
// 	type Thing struct {
// 		arg string
// 	}

// 	expected := []Thing{{arg: "a"}, {arg: "a"}}
// 	input := []Thing{{arg: "a"}, {arg: "b"}, {arg: "a"}}

// 	actual := slices.All(input, Thing{arg: "a"})

// 	assert.Equal(t, expected, actual)
// }

// func Test_AllFn_UserDefined(t *testing.T) {
// 	type Thing struct {
// 		arg string
// 	}

// 	expected := []Thing{{arg: "a"}, {arg: "a"}}
// 	input := []Thing{{arg: "a"}, {arg: "b"}, {arg: "a"}}

// 	actual := slices.AllFn(input, func(v Thing) bool {
// 		return v.arg == "a"
// 	})

// 	assert.Equal(t, expected, actual)
// }
