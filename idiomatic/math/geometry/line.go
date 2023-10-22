package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/math/types"
)

func NewLineSegmentXY[T types.Numbers](x1, y1, x2, y2 T) LineSegment[T] {
	return NewLineSegment[T](
		CartesianCoordinate[T]{X: x1, Y: y2},
		CartesianCoordinate[T]{X: x2, Y: y2},
	)
}

func NewLineSegment[T types.Numbers](start, end CartesianCoordinate[T]) LineSegment[T] {
	return LineSegment[T]{
		Start: start,
		End:   end,
	}
}

type LineSegment[T types.Numbers] struct {
	Start CartesianCoordinate[T]
	End   CartesianCoordinate[T]
}

func (t LineSegment[T]) Length() T {
	fn := func(x1, y1, x2, y2 float64) float64 {
		return math.Sqrt(math.Square((x2 - x1)) + math.Square((y2 - y1)))
	}

	return internal.QuadrupleToSingle(
		t.Start.X, t.Start.Y,
		t.End.X, t.End.Y,
		fn,
	)
}

func (t LineSegment[T]) PointAtPercent(percentage T) CartesianCoordinate[T] {
	fn := func(p1, p2, percentage float64) float64 {
		return p1 + p2*percentage/100
	}

	return CartesianCoordinate[T]{
		X: internal.TripleToSingle(t.Start.X, t.End.X, percentage, fn),
		Y: internal.TripleToSingle(t.Start.Y, t.End.Y, percentage, fn),
	}
}

func (t LineSegment[T]) Midpoint() CartesianCoordinate[T] {
	return t.PointAtPercent(50)
}

func (t LineSegment[T]) Slope() T {
	fn := func(x1, y1, x2, y2 float64) float64 {
		return (y2 - y1) / (x2 - x1)
	}

	return internal.QuadrupleToSingle(t.Start.X, t.Start.Y, t.End.X, t.End.Y, fn)
}

// https://martin-thoma.com/how-to-check-if-two-line-segments-intersect/
func (t LineSegment[T]) IntersectAt(line LineSegment[T]) (CartesianCoordinate[T], bool) {
	var coord CartesianCoordinate[T]
	var ok bool

	return coord, ok
}

func (t LineSegment[T]) Intersect(line LineSegment[T]) bool {
	var result bool

	o1 := lineSegmentOrientation(t, line.Start)
	o2 := lineSegmentOrientation(t, line.End)
	o3 := lineSegmentOrientation(line, t.Start)
	o4 := lineSegmentOrientation(line, t.End)

	if o1 != o2 && o3 != o4 {
		result = true
	}

	return result
}

var EPS = 1e-7

func lineSegmentPointOnLine[T types.Numbers](line LineSegment[T], coord CartesianCoordinate[T]) bool {
	var result bool

	return result
}

func lineSegmentOrientation[T types.Numbers](line LineSegment[T], coord CartesianCoordinate[T]) int {
	val := (line.End.Y-line.Start.Y)*(coord.X-line.End.X) - (line.End.X-line.Start.X)*(coord.X-line.End.Y)
	var result int

	if math.Abs(val) < 0 {
		return result
	} else if val > 0 {
		result = -1
	} else {
		result = 1
	}

	return result
}
