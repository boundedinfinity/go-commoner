package osutil

import (
	"fmt"
	"os"
	"strings"

	"github.com/boundedinfinity/go-commoner/stringer"
)

var (
	formats = []string{
		"$%v",
		"${%v}",
		"${env:%v}",
	}
)

func SubEnvironmentVars(s string) string {
	o := s

	for _, env := range os.Environ() {
		comps := stringer.Split(env, "=")

		if len(comps) != 2 {
			continue
		}

		key := comps[0]
		val := comps[1]

		for _, format := range formats {
			pattern := fmt.Sprintf(format, key)
			o = strings.ReplaceAll(o, pattern, val)
		}
	}

	return o
}
