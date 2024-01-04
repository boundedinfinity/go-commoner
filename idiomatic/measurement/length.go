package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"

func Parse(s string) (Length, error) {
	var length Length

	// var unit rational.Rational

	// for _,

	return length, nil
}

type Length struct {
	Number rational.Rational     `json:"unit,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Length) Unit() {

}

func (t Length) Metric() Length {
	switch t.System {
	case MeasurementSystems.Imperial:

		return Length{}
	default:
		return t
	}
}

func (t Length) Imperial() Length {
	return Length{}
}
