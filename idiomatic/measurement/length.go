package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/math/rational"

func Parse(s string) (Length, error) {
	var length Length

	return length, nil
}

type Length struct {
	Unit   rational.Rational     `json:"unit,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Length) Metric() Length {
	return Length{}
}

func (t Length) Imperial() Length {
	return Length{}
}
