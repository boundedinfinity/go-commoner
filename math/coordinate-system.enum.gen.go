//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package math

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

// /////////////////////////////////////////////////////////////////
//  AngleType Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t AngleType) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  AngleType JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AngleType) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *AngleType) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, AngleTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  AngleType YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AngleType) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *AngleType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, AngleTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  AngleType SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AngleType) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *AngleType) Scan(value interface{}) error {
	return enumer.Scan(value, t, AngleTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

var AngleTypes = struct {
	angleTypes
	Err    error
	Values []AngleType
}{
	angleTypes: angleTypes{
		Degrees: AngleType("Degrees"),
		Radians: AngleType("Radians"),
	},
	Err: fmt.Errorf("invalid AngleType"),
}

func init() {
	AngleTypes.Values = []AngleType{
		AngleTypes.Degrees,
		AngleTypes.Radians,
	}
}

func (t angleTypes) newErr(a any, values ...AngleType) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		AngleTypes.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t angleTypes) ParseFrom(v string, values ...AngleType) (AngleType, error) {
	var found AngleType
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, AngleType](v)(value) {
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

func (t angleTypes) Parse(v string) (AngleType, error) {
	return t.ParseFrom(v, AngleTypes.Values...)
}

func (t angleTypes) IsFrom(v string, values ...AngleType) bool {
	for _, value := range values {
		if enumer.IsEq[string, AngleType](v)(value) {
			return true
		}
	}
	return false
}

func (t angleTypes) Is(v string) bool {
	return t.IsFrom(v, AngleTypes.Values...)
}
