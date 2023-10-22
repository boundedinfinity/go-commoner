package types

import "golang.org/x/exp/constraints"

type Numbers interface {
	constraints.Float | constraints.Integer
}
