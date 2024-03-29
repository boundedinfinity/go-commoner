//////////////////////////////////////////////////////////////////
///                                                              /
///                          DO NOT EDIT                         /
///                                                              /
///              Manual changes will be overwritten.             /
///                                                              /
///        Generated by github.com/boundedinfinity/enumer        /
///                                                              /
//////////////////////////////////////////////////////////////////

package geometry

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

//////////////////////////////////////////////////////////////////
///                                                              /
///                             Type                             /
///                                                              /
//////////////////////////////////////////////////////////////////

type AngleType string

//////////////////////////////////////////////////////////////////
///                                                              /
///                    Stringer implemenation                    /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t AngleType) String() string {
	return string(t)
}

//////////////////////////////////////////////////////////////////
///                                                              /
///             JSON marshal/unmarshal implemenation             /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t AngleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *AngleType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	found, err := AngleTypes.Parse(s)

	if err != nil {
		return err
	}

	*t = found
	return nil
}

//////////////////////////////////////////////////////////////////
///                                                              /
///             YAML marshal/unmarshal implemenation             /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t AngleType) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *AngleType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	if err := unmarshal(&s); err != nil {
		return err
	}

	found, err := AngleTypes.Parse(s)

	if err != nil {
		return err
	}

	*t = found
	return nil
}

//////////////////////////////////////////////////////////////////
///                                                              /
///              XML marshal/unmarshal implemenation             /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t AngleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(string(t), start)
}

func (t *AngleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	found, err := AngleTypes.Parse(s)

	if err != nil {
		return err
	}

	*t = found
	return nil
}

//////////////////////////////////////////////////////////////////
///                                                              /
///              SQL marshal/unmarshal implemenation             /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t AngleType) Value() (driver.Value, error) {
	return string(t), nil
}

func (t *AngleType) Scan(value interface{}) error {
	if value == nil {
		return AngleTypes.errf(value)
	}

	dv, err := driver.String.ConvertValue(value)

	if err != nil {
		return err
	}

	s, ok := dv.(string)

	if !ok {
		return AngleTypes.errf(value)
	}

	found, err := AngleTypes.Parse(s)

	if err != nil {
		return err
	}

	*t = found
	return nil
}

//////////////////////////////////////////////////////////////////
///                                                              /
///                       Companion struct                       /
///                                                              /
//////////////////////////////////////////////////////////////////

var AngleTypes = angleTypes{
	Degrees: AngleType("degrees"),
	Err:     fmt.Errorf("invalid AngleType"),
	Radians: AngleType("radians"),
}

type angleTypes struct {
	Err      error
	errf     func(any, ...AngleType) error
	parseMap map[AngleType][]string
	Degrees  AngleType
	Radians  AngleType
}

func (t angleTypes) Values() []AngleType {
	return []AngleType{
		AngleTypes.Degrees,
		AngleTypes.Radians,
	}
}

func (t angleTypes) ParseFrom(v string, items ...AngleType) (AngleType, error) {
	var found AngleType
	var ok bool

	for _, item := range items {
		matchers, ok2 := t.parseMap[item]

		if !ok2 {
			continue
		}

		for _, matcher := range matchers {
			if v == matcher {
				found = item
				ok = true
				break
			}
		}

		if !ok {
			return found, t.errf(v, items...)
		}

		return found, nil
	}

	return found, nil
}

func (t angleTypes) Parse(v string) (AngleType, error) {
	return t.ParseFrom(v, t.Values()...)
}

func (t angleTypes) IsFrom(v string, items ...AngleType) bool {
	_, err := t.ParseFrom(v, items...)
	return err == nil
}

func (t angleTypes) Is(v string) bool {
	return t.IsFrom(v, t.Values()...)
}

//////////////////////////////////////////////////////////////////
///                                                              /
///                        Initialization                        /
///                                                              /
//////////////////////////////////////////////////////////////////

func init() {
	AngleTypes.parseMap = map[AngleType][]string{
		AngleTypes.Degrees: {"degrees", "Degrees"},
		AngleTypes.Radians: {"radians", "Radians"},
	}

	AngleTypes.errf = func(v any, items ...AngleType) error {
		var xs []string

		for _, item := range items {
			if x, ok := AngleTypes.parseMap[item]; ok {
				xs = append(xs, x...)
			}
		}

		return fmt.Errorf(
			"%w: %v is not one of %s",
			AngleTypes.Err,
			v,
			strings.Join(xs, ","),
		)
	}
}
