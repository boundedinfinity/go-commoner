//////////////////////////////////////////////////////////////////
///                                                              /
///                          DO NOT EDIT                         /
///                                                              /
///              Manual changes will be overwritten.             /
///                                                              /
///        Generated by github.com/boundedinfinity/enumer        /
///                                                              /
//////////////////////////////////////////////////////////////////

package measurement

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

type ImperialMassUnit string

//////////////////////////////////////////////////////////////////
///                                                              /
///                    Stringer implemenation                    /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t ImperialMassUnit) String() string {
	return string(t)
}

//////////////////////////////////////////////////////////////////
///                                                              /
///             JSON marshal/unmarshal implemenation             /
///                                                              /
//////////////////////////////////////////////////////////////////

func (t ImperialMassUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *ImperialMassUnit) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	found, err := ImperialMassUnits.Parse(s)

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

func (t ImperialMassUnit) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *ImperialMassUnit) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	if err := unmarshal(&s); err != nil {
		return err
	}

	found, err := ImperialMassUnits.Parse(s)

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

func (t ImperialMassUnit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(string(t), start)
}

func (t *ImperialMassUnit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	found, err := ImperialMassUnits.Parse(s)

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

func (t ImperialMassUnit) Value() (driver.Value, error) {
	return string(t), nil
}

func (t *ImperialMassUnit) Scan(value interface{}) error {
	if value == nil {
		return ImperialMassUnits.errf(value)
	}

	dv, err := driver.String.ConvertValue(value)

	if err != nil {
		return err
	}

	s, ok := dv.(string)

	if !ok {
		return ImperialMassUnits.errf(value)
	}

	found, err := ImperialMassUnits.Parse(s)

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

var ImperialMassUnits = imperialMassUnits{
	Err:               fmt.Errorf("invalid ImperialMassUnit"),
	Invalid:           ImperialMassUnit("invalid"),
	Grain:             ImperialMassUnit("grain"),
	PennyWeight:       ImperialMassUnit("penny-weight"),
	Scruple:           ImperialMassUnit("scruple"),
	Drachm:            ImperialMassUnit("drachm"),
	Ounce:             ImperialMassUnit("ounce"),
	OunceApothecaries: ImperialMassUnit("ounce-apothecaries"),
	Pound:             ImperialMassUnit("pound"),
	Dram:              ImperialMassUnit("dram"),
	Stone:             ImperialMassUnit("stone"),
	Quarter:           ImperialMassUnit("quarter"),
	Cental:            ImperialMassUnit("cental"),
	HundredWeight:     ImperialMassUnit("hundred-weight"),
	Ton:               ImperialMassUnit("ton"),
	MetricTon:         ImperialMassUnit("metric-ton"),
	Quintal:           ImperialMassUnit("quintal"),
	Slug:              ImperialMassUnit("slug"),
}

type imperialMassUnits struct {
	Err               error
	errf              func(any, ...ImperialMassUnit) error
	parseMap          map[ImperialMassUnit][]string
	Invalid           ImperialMassUnit
	Grain             ImperialMassUnit
	PennyWeight       ImperialMassUnit
	Scruple           ImperialMassUnit
	Drachm            ImperialMassUnit
	Ounce             ImperialMassUnit
	OunceApothecaries ImperialMassUnit
	Pound             ImperialMassUnit
	Dram              ImperialMassUnit
	Stone             ImperialMassUnit
	Quarter           ImperialMassUnit
	Cental            ImperialMassUnit
	HundredWeight     ImperialMassUnit
	Ton               ImperialMassUnit
	MetricTon         ImperialMassUnit
	Quintal           ImperialMassUnit
	Slug              ImperialMassUnit
}

func (t imperialMassUnits) Values() []ImperialMassUnit {
	return []ImperialMassUnit{
		ImperialMassUnits.Grain,
		ImperialMassUnits.PennyWeight,
		ImperialMassUnits.Scruple,
		ImperialMassUnits.Drachm,
		ImperialMassUnits.Ounce,
		ImperialMassUnits.OunceApothecaries,
		ImperialMassUnits.Pound,
		ImperialMassUnits.Dram,
		ImperialMassUnits.Stone,
		ImperialMassUnits.Quarter,
		ImperialMassUnits.Cental,
		ImperialMassUnits.HundredWeight,
		ImperialMassUnits.Ton,
		ImperialMassUnits.MetricTon,
		ImperialMassUnits.Quintal,
		ImperialMassUnits.Slug,
	}
}

func (t imperialMassUnits) ParseFrom(v string, items ...ImperialMassUnit) (ImperialMassUnit, error) {
	var found ImperialMassUnit
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

		if ok {
			break
		}
	}

	if !ok {
		return found, t.errf(v, items...)
	}

	return found, nil
}

func (t imperialMassUnits) Parse(v string) (ImperialMassUnit, error) {
	return t.ParseFrom(v, t.Values()...)
}

func (t imperialMassUnits) IsFrom(v string, items ...ImperialMassUnit) bool {
	_, err := t.ParseFrom(v, items...)
	return err == nil
}

func (t imperialMassUnits) Is(v string) bool {
	return t.IsFrom(v, t.Values()...)
}

//////////////////////////////////////////////////////////////////
///                                                              /
///                        Initialization                        /
///                                                              /
//////////////////////////////////////////////////////////////////

func init() {
	ImperialMassUnits.parseMap = map[ImperialMassUnit][]string{
		ImperialMassUnits.Cental:            {"cental", "Cental"},
		ImperialMassUnits.Drachm:            {"drachm", "Drachm"},
		ImperialMassUnits.Dram:              {"dram", "Dram"},
		ImperialMassUnits.Grain:             {"grain", "Grain"},
		ImperialMassUnits.HundredWeight:     {"hundred-weight", "HundredWeight"},
		ImperialMassUnits.MetricTon:         {"metric-ton", "MetricTon"},
		ImperialMassUnits.Ounce:             {"ounce", "Ounce"},
		ImperialMassUnits.OunceApothecaries: {"ounce-apothecaries", "OunceApothecaries"},
		ImperialMassUnits.PennyWeight:       {"penny-weight", "PennyWeight"},
		ImperialMassUnits.Pound:             {"pound", "Pound"},
		ImperialMassUnits.Quarter:           {"quarter", "Quarter"},
		ImperialMassUnits.Quintal:           {"quintal", "Quintal"},
		ImperialMassUnits.Scruple:           {"scruple", "Scruple"},
		ImperialMassUnits.Slug:              {"slug", "Slug"},
		ImperialMassUnits.Stone:             {"stone", "Stone"},
		ImperialMassUnits.Ton:               {"ton", "Ton"},
	}

	ImperialMassUnits.errf = func(v any, items ...ImperialMassUnit) error {
		var xs []string

		for _, item := range items {
			if x, ok := ImperialMassUnits.parseMap[item]; ok {
				xs = append(xs, x...)
			}
		}

		return fmt.Errorf(
			"%w: %v is not one of %s",
			ImperialMassUnits.Err,
			v,
			strings.Join(xs, ","),
		)
	}
}
