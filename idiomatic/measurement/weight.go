package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"

type Weight struct {
	Unit   rational.Rational     `json:"unit,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Weight) Metric() Weight {
	return Weight{}
}

func (t Weight) Imperial() Weight {
	return Weight{}
}
