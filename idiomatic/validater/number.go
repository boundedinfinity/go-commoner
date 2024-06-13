package validater

import "github.com/boundedinfinity/go-commoner/errorer"

var ErrNumberIsInvalid = errorer.New("number")

// Builder

// func Number[T constraints.Integer | constraints.Float]() *NumberValidatorBuilder[T] {
// 	return &NumberValidatorBuilder[T]{
// 		fns: make([]func(T) error, 0),
// 	}
// }

// type NumberValidatorBuilder[T constraints.Integer | constraints.Float] struct {
// 	fns []func(T) error
// }

// func (t *NumberValidatorBuilder[T]) Validator(elem T) Validater {
// 	return &validater[T]{
// 		fns:  t.fns,
// 		elem: elem,
// 	}
// }

// func (t *NumberValidatorBuilder[T]) Min(min T) *NumberValidatorBuilder[T] {
// 	t.fns = append(t.fns, NumberMin[T](min))
// 	return t
// }

// func (t *NumberValidatorBuilder[T]) Max(max T) *NumberValidatorBuilder[T] {
// 	t.fns = append(t.fns, NumberMax[T](max))
// 	return t
// }

// func (t *NumberValidatorBuilder[T]) NotOneOf(elems ...T) *NumberValidatorBuilder[T] {
// 	t.fns = append(t.fns, NotOneOf[T](elems...))
// 	return t
// }

// // Stand alone

// func NotOneOf[T constraints.Integer | constraints.Float](elems ...T) func(s T) error {
// 	return func(n T) error {
// 		for i, elem := range elems {
// 			if elem == n {
// 				return fmt.Errorf(
// 					"%w %v cannot be equal elems[%v] = %v",
// 					ErrNumberIsInvalid, n, i, elems[i],
// 				)
// 			}
// 		}

// 		return nil
// 	}
// }

// func NumberGreaterThan[T constraints.Integer | constraints.Float](min T) func(s T) error {
// 	return func(n T) error {
// 		if n <= min {
// 			return fmt.Errorf("%w less than %v", ErrNumberIsInvalid, min)
// 		}

// 		return nil
// 	}
// }

// func NumberMin[T constraints.Integer | constraints.Float](min T) func(s T) error {
// 	return func(n T) error {
// 		if n < min {
// 			return fmt.Errorf("%w less than %v", ErrNumberIsInvalid, min)
// 		}

// 		return nil
// 	}
// }

// func NumberMax[T constraints.Integer | constraints.Float](max T) func(s T) error {
// 	return func(n T) error {
// 		if n > max {
// 			return fmt.Errorf("%w greater than %v", ErrNumberIsInvalid, max)
// 		}

// 		return nil
// 	}
// }
