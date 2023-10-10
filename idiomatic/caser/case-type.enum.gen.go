//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package caser

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

// /////////////////////////////////////////////////////////////////
//  CaseType Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t CaseType) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  CaseType JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CaseType) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *CaseType) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, CaseTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  CaseType YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CaseType) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *CaseType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, CaseTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//  CaseType SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t CaseType) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *CaseType) Scan(value interface{}) error {
	return enumer.Scan(value, t, CaseTypes.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

var CaseTypes = struct {
	caseTypes
	Err    error
	Values []CaseType
}{
	caseTypes: caseTypes{
		Camel:      CaseType("Camel"),
		Kebab:      CaseType("Kebab"),
		Kebabupper: CaseType("Kebabupper"),
		Pascal:     CaseType("Pascal"),
		Phrase:     CaseType("Phrase"),
		Snake:      CaseType("Snake"),
		Snakeupper: CaseType("Snakeupper"),
		Unknown:    CaseType("Unknown"),
	},
	Err: fmt.Errorf("invalid CaseType"),
}

func init() {
	CaseTypes.Values = []CaseType{
		CaseTypes.Camel,
		CaseTypes.Kebab,
		CaseTypes.Kebabupper,
		CaseTypes.Pascal,
		CaseTypes.Phrase,
		CaseTypes.Snake,
		CaseTypes.Snakeupper,
		CaseTypes.Unknown,
	}
}

func (t caseTypes) newErr(a any, values ...CaseType) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		CaseTypes.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t caseTypes) ParseFrom(v string, values ...CaseType) (CaseType, error) {
	var found CaseType
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, CaseType](v)(value) {
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

func (t caseTypes) Parse(v string) (CaseType, error) {
	return t.ParseFrom(v, CaseTypes.Values...)
}

func (t caseTypes) IsFrom(v string, values ...CaseType) bool {
	for _, value := range values {
		if enumer.IsEq[string, CaseType](v)(value) {
			return true
		}
	}
	return false
}

func (t caseTypes) Is(v string) bool {
	return t.IsFrom(v, CaseTypes.Values...)
}
