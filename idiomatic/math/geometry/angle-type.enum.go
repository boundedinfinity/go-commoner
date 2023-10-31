package geometry

//go:generate enumer -path=./angle-type.enum.go

type AngleType string

type angleTypes struct {
	Degrees AngleType `enum:"degress"`
	Radians AngleType `enum:"radians"`
}
