package geometry

import "golang.org/x/exp/constraints"

type geometryNumber interface {
	constraints.Float | ~int16 | ~int32 | ~int64 | ~uint16 | ~uint32 | ~uint64
}
