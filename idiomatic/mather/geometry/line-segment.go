package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

func NewLineSegmentXY[T geometryNumber](x1, y1, x2, y2 T) LineSegment[T] {
	return NewLineSegmentCoords[T](
		CartesianCoordinate[T]{X: x1, Y: y2},
		CartesianCoordinate[T]{X: x2, Y: y2},
	)
}

func NewLineSegmentCoords[T geometryNumber](start, end CartesianCoordinate[T]) LineSegment[T] {
	return LineSegment[T]{
		Start: start,
		End:   end,
	}
}

type LineSegment[T geometryNumber] struct {
	Start CartesianCoordinate[T]
	End   CartesianCoordinate[T]
}

func (t LineSegment[T]) Length() T {
	fn := func(x1, y1, x2, y2 float64) float64 {
		return mather.Sqrt(mather.Square((x2 - x1)) + mather.Square((y2 - y1)))
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

func (t LineSegment[T]) SlopeIntercept() SlopeInterceptLine[T] {
	slope := t.Slope()

	return SlopeInterceptLine[T]{
		M: slope,
		B: t.Start.Y / slope * t.Start.X,
	}
}

func (t LineSegment[T]) Standard() StandardLine[T] {
	sif := t.SlopeIntercept()

	return StandardLine[T]{
		A: -sif.M,
		B: 1,
		C: -sif.B,
	}
}

// https://www.cuemath.com/geometry/intersection-of-two-lines/
// https://www.cuemath.com/algebra/cross-multiplication-method/
func (t LineSegment[T]) IntersectAt(segment LineSegment[T]) CartesianCoordinate[T] {
	l1 := t.Standard()
	l2 := segment.Standard()

	return CartesianCoordinate[T]{
		X: (l1.B*l2.C - l2.B*l1.C) / (l1.A*l2.B - l2.A*l1.B),
		Y: (l1.C*l2.A - l2.A*l1.A) / (l1.A*l2.B - l2.A*l1.B),
	}
}
