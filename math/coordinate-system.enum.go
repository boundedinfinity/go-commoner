package math

//go:generate enumer -path=./coordinate-system.enum.go

type AngleType string

type angleTypes struct {
	Degrees AngleType
	Radians AngleType
}
