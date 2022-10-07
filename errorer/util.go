package errorer

import "fmt"

func Errorv(err error, v interface{}) error {
	return fmt.Errorf("%w : %v", err, v)
}

func Errorfn(err error) func(v interface{}) error {
	return func(v interface{}) error {
		return Errorv(err, v)
	}
}
