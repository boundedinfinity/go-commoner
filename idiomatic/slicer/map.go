package slicer

import (
	"errors"
	"fmt"
)

func Map[T any, R any](fn func(T) R, elems ...T) []R {
	var results []R

	for _, elem := range elems {
		results = append(results, fn(elem))
	}

	return results
}

func MapI[T any, R any](fn func(int, T) R, elems ...T) []R {
	var results []R

	for i, elem := range elems {
		results = append(results, fn(i, elem))
	}

	return results
}

func MapErr[T any, R any](fn func(T) (R, error), elems ...T) ([]R, error) {
	var results []R

	for i, elem := range elems {
		result, err := fn(elem)

		if err != nil {
			return results, &errMapFuncDetails{Index: i, Err: err}
		}

		results = append(results, result)
	}

	return results, nil
}

func MapErrI[T any, R any](fn func(int, T) (R, error), elems ...T) ([]R, error) {
	var results []R

	for i, elem := range elems {
		result, err := fn(i, elem)

		if err != nil {
			return results, &errMapFuncDetails{Index: i, Err: err}
		}

		results = append(results, result)
	}

	return results, nil
}

func MapChain[T any, R any](fns []func(T) R, elems ...T) []R {
	var results []R

	for _, fn := range fns {
		results = Map(fn, elems...)
	}

	return results
}

func MapChainI[T any, R any](fns []func(int, T) R, elems ...T) []R {
	var results []R

	for _, fn := range fns {
		results = MapI(fn, elems...)
	}

	return results
}

func MapChainErr[T any, R any](fns []func(T) (R, error), elems ...T) ([]R, error) {
	var results []R
	var err error

	for _, fn := range fns {
		results, err = MapErr(fn, elems...)

		if err != nil {
			break
		}
	}

	return results, err
}

func MapChainErrI[T any, R any](fns []func(int, T) (R, error), elems ...T) ([]R, error) {
	var results []R

	for i, fn := range fns {
		intermediate, err := MapErrI(fn, elems...)

		if err != nil {
			return results, &errMapChainFuncDetails{
				ElemIndex: err.(*errMapFuncDetails).Index,
				FuncIndex: i,
				Err:       err.(*errMapFuncDetails).Err,
			}
		}

		results = intermediate
	}

	return results, nil
}

var ErrMapFunc = errors.New("map err")

type errMapFuncDetails struct {
	Index int
	Err   error
}

func (this errMapFuncDetails) Error() string {
	return fmt.Sprintf("%v[%d] : %s", ErrMapFunc.Error(), this.Index, this.Err.Error())
}

func (this errMapFuncDetails) Unwrap() error {
	return ErrMapFunc
}

var ErrMapChanFunc = errors.New("map err")

type errMapChainFuncDetails struct {
	ElemIndex int
	FuncIndex int
	Err       error
}

func (this errMapChainFuncDetails) Error() string {
	return fmt.Sprintf("%v[%d] : %s", ErrMapFunc.Error(), this.ElemIndex, this.Err.Error())
}

func (this errMapChainFuncDetails) Unwrap() error {
	return ErrMapChanFunc
}
