package errorer

import (
	"fmt"
)

func Errorf(format string, a ...any) Error {
	return Err(fmt.Errorf(format, a...))
}

func Err(err error) Error {
	return Error{
		err: err,
	}
}

func None() Error {
	return Error{
		err: nil,
	}
}

type Error struct {
	err error
}

func (e Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}
