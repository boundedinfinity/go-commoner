package types

import "golang.org/x/exp/constraints"

type Float constraints.Float
type Integer constraints.Integer
type Ordered constraints.Ordered

type Number interface {
	Integer | Float
}
