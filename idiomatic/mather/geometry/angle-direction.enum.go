package geometry

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	enumer "github.com/boundedinfinity/enumer"
)

//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

type AngleDirection string

func (t AngleDirection) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  JSON serialization
// /////////////////////////////////////////////////////////////////

func (t AngleDirection) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *AngleDirection) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, AngleDirections.Parse)
}

// /////////////////////////////////////////////////////////////////
//  YAML serialization
// /////////////////////////////////////////////////////////////////

func (t AngleDirection) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *AngleDirection) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, AngleDirections.Parse)
}

// /////////////////////////////////////////////////////////////////
//  XML serialization
// /////////////////////////////////////////////////////////////////

func (t AngleDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return enumer.MarshalXML(t, e, start)
}

func (t *AngleDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return enumer.UnmarshalXML(t, AngleDirections.Parse, d, start)
}

// /////////////////////////////////////////////////////////////////
//  SQL serialization
// /////////////////////////////////////////////////////////////////

func (t AngleDirection) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *AngleDirection) Scan(value interface{}) error {
	return enumer.Scan(value, t, AngleDirections.Parse)
}

// /////////////////////////////////////////////////////////////////
//  Companion
// /////////////////////////////////////////////////////////////////

type angleDirections struct {
	Clockwise        AngleDirection
	CounterClockwise AngleDirection
	Values           []AngleDirection
	Err              error
}

var AngleDirections = angleDirections{

	Clockwise:        AngleDirection("clockwise"),
	CounterClockwise: AngleDirection("counter-clockwise"),
	Err:              fmt.Errorf("invalid AngleDirection"),
}

func init() {
	AngleDirections.Values = []AngleDirection{
		AngleDirections.Clockwise,
		AngleDirections.CounterClockwise,
	}
}

func (t angleDirections) newErr(a any, values ...AngleDirection) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		AngleDirections.Err,
		a,
		enumer.Join(values, ", "))
}

func (t angleDirections) ParseFrom(v string, values ...AngleDirection) (AngleDirection, error) {
	var found AngleDirection
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, AngleDirection](v)(value) {
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

func (t angleDirections) Parse(v string) (AngleDirection, error) {
	return t.ParseFrom(v, AngleDirections.Values...)
}

func (t angleDirections) IsFrom(v string, values ...AngleDirection) bool {
	for _, value := range values {
		if enumer.IsEq[string, AngleDirection](v)(value) {
			return true
		}
	}
	return false
}

func (t angleDirections) Is(v string) bool {
	return t.IsFrom(v, AngleDirections.Values...)
}
