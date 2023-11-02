//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package measurement

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

// /////////////////////////////////////////////////////////////////
//  MetricUnit Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t MetricUnit) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  MetricUnit JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MetricUnit) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *MetricUnit) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, MetricUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//  MetricUnit YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MetricUnit) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *MetricUnit) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, MetricUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//  MetricUnit SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MetricUnit) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *MetricUnit) Scan(value interface{}) error {
	return enumer.Scan(value, t, MetricUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

var MetricUnits = struct {
	metricUnits
	Err    error
	Values []MetricUnit
}{
	metricUnits: metricUnits{
		Tera:  MetricUnit("tera"),
		Giga:  MetricUnit("giga"),
		Mega:  MetricUnit("mega"),
		Kilo:  MetricUnit("kilo"),
		Hecto: MetricUnit("hecto"),
		Deca:  MetricUnit("deca"),
		Unit:  MetricUnit("unit"),
		Deci:  MetricUnit("deci"),
		Centi: MetricUnit("centi"),
		Milli: MetricUnit("milli"),
		Micro: MetricUnit("micro"),
		Nano:  MetricUnit("nano"),
		Pico:  MetricUnit("pico"),
	},
	Err: fmt.Errorf("invalid MetricUnit"),
}

func init() {
	MetricUnits.Values = []MetricUnit{
		MetricUnits.Tera,
		MetricUnits.Giga,
		MetricUnits.Mega,
		MetricUnits.Kilo,
		MetricUnits.Hecto,
		MetricUnits.Deca,
		MetricUnits.Unit,
		MetricUnits.Deci,
		MetricUnits.Centi,
		MetricUnits.Milli,
		MetricUnits.Micro,
		MetricUnits.Nano,
		MetricUnits.Pico,
	}
}

func (t metricUnits) newErr(a any, values ...MetricUnit) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		MetricUnits.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t metricUnits) ParseFrom(v string, values ...MetricUnit) (MetricUnit, error) {
	var found MetricUnit
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, MetricUnit](v)(value) {
			found = value
			ok = true
			break
		}
	}

	if !ok {
		return found, t.newErr(v, values...)
	}

	return found, nil
}

func (t metricUnits) Parse(v string) (MetricUnit, error) {
	return t.ParseFrom(v, MetricUnits.Values...)
}

func (t metricUnits) IsFrom(v string, values ...MetricUnit) bool {
	for _, value := range values {
		if enumer.IsEq[string, MetricUnit](v)(value) {
			return true
		}
	}
	return false
}

func (t metricUnits) Is(v string) bool {
	return t.IsFrom(v, MetricUnits.Values...)
}