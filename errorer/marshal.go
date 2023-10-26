package errorer

import (
	"encoding/json"
	"errors"
)

func (e Errorer) MarshalJSON() ([]byte, error) {
	if e.wrapped == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(e.Error())
}

func (e *Errorer) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e.wrapped = errors.New(s)

	return nil
}
