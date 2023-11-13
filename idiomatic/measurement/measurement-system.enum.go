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

type MeasurementSystem string

// /////////////////////////////////////////////////////////////////
//  MeasurementSystem Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t MeasurementSystem) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  MeasurementSystem JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MeasurementSystem) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *MeasurementSystem) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, MeasurementSystems.Parse)
}

// /////////////////////////////////////////////////////////////////
//  MeasurementSystem YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MeasurementSystem) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *MeasurementSystem) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, MeasurementSystems.Parse)
}

// /////////////////////////////////////////////////////////////////
//  MeasurementSystem SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t MeasurementSystem) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *MeasurementSystem) Scan(value interface{}) error {
	return enumer.Scan(value, t, MeasurementSystems.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type measurementSystems struct {
	Metric   MeasurementSystem
	Imperial MeasurementSystem
	Values   []MeasurementSystem
	Err      error
}

var MeasurementSystems = measurementSystems{
	Metric:   MeasurementSystem("metric"),
	Imperial: MeasurementSystem("imperial"),
	Err:      fmt.Errorf("invalid MeasurementSystem"),
}

func init() {
	MeasurementSystems.Values = []MeasurementSystem{
		MeasurementSystems.Metric,
		MeasurementSystems.Imperial,
	}
}

func (t measurementSystems) newErr(a any, values ...MeasurementSystem) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		MeasurementSystems.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t measurementSystems) ParseFrom(v string, values ...MeasurementSystem) (MeasurementSystem, error) {
	var found MeasurementSystem
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, MeasurementSystem](v)(value) {
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

func (t measurementSystems) Parse(v string) (MeasurementSystem, error) {
	return t.ParseFrom(v, MeasurementSystems.Values...)
}

func (t measurementSystems) IsFrom(v string, values ...MeasurementSystem) bool {
	for _, value := range values {
		if enumer.IsEq[string, MeasurementSystem](v)(value) {
			return true
		}
	}
	return false
}

func (t measurementSystems) Is(v string) bool {
	return t.IsFrom(v, MeasurementSystems.Values...)
}
