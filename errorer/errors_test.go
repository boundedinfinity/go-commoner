package errorer_test

import (
	"fmt"

	"testing"

	"github.com/boundedinfinity/bounded-avatar/model/util/errorer"
)

func Test_error(t *testing.T) {
	ErrRoot1 := errorer.New("root error 1")
	ErrRoot2 := errorer.New("root error 2")
	errRootFn := errorer.Func(ErrRoot1, ErrRoot2)

	err1 := errRootFn("something when wrong! %s", "hmm...")
	s := err1.Error()

	fmt.Println(s)
}
