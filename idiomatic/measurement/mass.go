package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"

type Mass struct {
	Value  rational.Rational     `json:"value,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Unit   ImperialMassUnit      `json:"unit,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Mass) Metric() Mass {
	return Mass{}
}

func (t Mass) Imperial() Mass {
	return Mass{}
}
