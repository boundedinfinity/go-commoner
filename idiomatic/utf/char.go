package utf

import (
	"encoding/json"
	"fmt"
)

type UtfChar byte

func (t UtfChar) HtmlNumber() string {
	return fmt.Sprintf("&#%02d;", t)
}

func (t UtfChar) String() string {
	return string(t)
}

func (t UtfChar) MarshalJSON() ([]byte, error) {
	return json.Marshal(byte(t))
}

func (t *UtfChar) UnmarshalJSON(data []byte) error {
	var v byte

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Utf7.Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t UtfChar) MarshalYAML() (interface{}, error) {
	return byte(t), nil
}

func (t *UtfChar) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v byte

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Utf7.Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
