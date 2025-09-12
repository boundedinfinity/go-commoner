package mapper

import (
	"errors"
	"fmt"
)

var ErrMapper = errors.New("mapper error")

type MapperErrDetails[K comparable, V any] struct {
	Key    K
	Val    V
	Reason error
}

func (e MapperErrDetails[K, V]) Error() string {
	return fmt.Sprintf(
		"%s : [%v, %+v] : %s",
		ErrMapper.Error(),
		e.Key, e.Val, e.Reason.Error(),
	)
}

func (e MapperErrDetails[K, V]) Unwrap() []error {
	return []error{ErrMapper, e.Reason}
}
