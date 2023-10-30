package validator

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

// func (t *NumberValidatorBuilder[T]) Validator(item T) Validater {
// 	return &validater[T]{
// 		fns:  t.fns,
// 		item: item,
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

// func (t *NumberValidatorBuilder[T]) NotOneOf(items ...T) *NumberValidatorBuilder[T] {
// 	t.fns = append(t.fns, NotOneOf[T](items...))
// 	return t
// }

// // Stand alone

// func NotOneOf[T constraints.Integer | constraints.Float](items ...T) func(s T) error {
// 	return func(n T) error {
// 		for i, item := range items {
// 			if item == n {
// 				return fmt.Errorf(
// 					"%w %v cannot be equal items[%v] = %v",
// 					ErrNumberIsInvalid, n, i, items[i],
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
