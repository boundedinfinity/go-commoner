package errorer

import (
	"encoding/json"
	"errors"
)

func (e Error) MarshalJSON() ([]byte, error) {
	if e.err != nil {
		return json.Marshal(e.Error())
	}

	return json.Marshal(nil)
}

func (e *Error) UnmarshalJSON(data []byte) error {
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

	e.err = errors.New(s)

	return nil
}
