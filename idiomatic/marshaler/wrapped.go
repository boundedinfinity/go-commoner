package marshaler

import (
	"encoding/json"
	"reflect"
)

type marshalerContext struct {
	typ reflect.Type
}

type marshalWrapper struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type unmarshalWrapper struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}
