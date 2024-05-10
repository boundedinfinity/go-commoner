package util

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func Dumpj(v interface{}) string {
	bs, _ := json.MarshalIndent(v, "", "    ")
	return string(bs)
}

func Dumpy(v interface{}) string {
	bs, _ := yaml.Marshal(v)
	return string(bs)
}
