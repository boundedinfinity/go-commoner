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

type ImperialLengthUnit string

// /////////////////////////////////////////////////////////////////
//  ImperialLengthUnit Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t ImperialLengthUnit) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  ImperialLengthUnit JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t ImperialLengthUnit) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *ImperialLengthUnit) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, ImperialLengthUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//  ImperialLengthUnit YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t ImperialLengthUnit) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *ImperialLengthUnit) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, ImperialLengthUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//  ImperialLengthUnit SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t ImperialLengthUnit) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *ImperialLengthUnit) Scan(value interface{}) error {
	return enumer.Scan(value, t, ImperialLengthUnits.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type imperialLengthUnits struct {
	Twip         ImperialLengthUnit
	Thou         ImperialLengthUnit
	Barleycorn   ImperialLengthUnit
	Inch         ImperialLengthUnit
	Hand         ImperialLengthUnit
	Foot         ImperialLengthUnit
	Yard         ImperialLengthUnit
	Chain        ImperialLengthUnit
	Furlong      ImperialLengthUnit
	Mile         ImperialLengthUnit
	League       ImperialLengthUnit
	Fathom       ImperialLengthUnit
	Cable        ImperialLengthUnit
	NauticalMile ImperialLengthUnit
	Link         ImperialLengthUnit
	Rod          ImperialLengthUnit
	Values       []ImperialLengthUnit
	Err          error
}

var ImperialLengthUnits = imperialLengthUnits{
	Twip:         ImperialLengthUnit("twip"),
	Thou:         ImperialLengthUnit("thou"),
	Barleycorn:   ImperialLengthUnit("barleycorn"),
	Inch:         ImperialLengthUnit("inch"),
	Hand:         ImperialLengthUnit("hand"),
	Foot:         ImperialLengthUnit("foot"),
	Yard:         ImperialLengthUnit("yard"),
	Chain:        ImperialLengthUnit("chain"),
	Furlong:      ImperialLengthUnit("furlong"),
	Mile:         ImperialLengthUnit("mile"),
	League:       ImperialLengthUnit("league"),
	Fathom:       ImperialLengthUnit("fathom"),
	Cable:        ImperialLengthUnit("cable"),
	NauticalMile: ImperialLengthUnit("nautical-mile"),
	Link:         ImperialLengthUnit("link"),
	Rod:          ImperialLengthUnit("rod"),
	Err:          fmt.Errorf("invalid ImperialLengthUnit"),
}

func init() {
	ImperialLengthUnits.Values = []ImperialLengthUnit{
		ImperialLengthUnits.Twip,
		ImperialLengthUnits.Thou,
		ImperialLengthUnits.Barleycorn,
		ImperialLengthUnits.Inch,
		ImperialLengthUnits.Hand,
		ImperialLengthUnits.Foot,
		ImperialLengthUnits.Yard,
		ImperialLengthUnits.Chain,
		ImperialLengthUnits.Furlong,
		ImperialLengthUnits.Mile,
		ImperialLengthUnits.League,
		ImperialLengthUnits.Fathom,
		ImperialLengthUnits.Cable,
		ImperialLengthUnits.NauticalMile,
		ImperialLengthUnits.Link,
		ImperialLengthUnits.Rod,
	}
}

func (t imperialLengthUnits) newErr(a any, values ...ImperialLengthUnit) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		ImperialLengthUnits.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t imperialLengthUnits) ParseFrom(v string, values ...ImperialLengthUnit) (ImperialLengthUnit, error) {
	var found ImperialLengthUnit
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, ImperialLengthUnit](v)(value) {
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

func (t imperialLengthUnits) Parse(v string) (ImperialLengthUnit, error) {
	return t.ParseFrom(v, ImperialLengthUnits.Values...)
}

func (t imperialLengthUnits) IsFrom(v string, values ...ImperialLengthUnit) bool {
	for _, value := range values {
		if enumer.IsEq[string, ImperialLengthUnit](v)(value) {
			return true
		}
	}
	return false
}

func (t imperialLengthUnits) Is(v string) bool {
	return t.IsFrom(v, ImperialLengthUnits.Values...)
}
