package geometry

//go:generate enumer -path=./angle-direction.enum.go

type AngleDirection string

type angleDirections struct {
	Clockwise        AngleDirection
	CounterClockwise AngleDirection
}
