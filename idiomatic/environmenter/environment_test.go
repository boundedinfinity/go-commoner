package environmenter_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/environmenter"
	"github.com/boundedinfinity/go-commoner/idiomatic/sheller"
	"github.com/stretchr/testify/assert"
)

func Test_Environmenter_WithLines(t *testing.T) {
	tcs := []struct {
		name        string
		environment *environmenter.Environment
		inputs      []string
		expected    []string
	}{
		{
			name: "sub with defaults",
			environment: environmenter.Shell(
				sheller.Shells.Bash,
				environmenter.KV("THING=/thing"),
			),
			inputs: []string{
				"$THING/dir1",
				"${THING}/dir2",
			},
			expected: []string{
				"/thing/dir1",
				"/thing/dir2",
			},
		},
		{
			name: "sub with embedded variable substitution",
			environment: environmenter.Shell(
				sheller.Shells.Bash,
				environmenter.KV(
					"THING=/thing",
					"OTHER=$THING/other",
				),
			),
			inputs: []string{
				"$OTHER/dir1",
				"${OTHER}/dir2",
			},
			expected: []string{
				"/thing/other/dir1",
				"/thing/other/dir2",
			},
		},
		{
			name: "XDG dirs and HOME",
			environment: environmenter.Shell(
				sheller.Shells.Bash,
				environmenter.XdgDirs(),
				environmenter.KV("HOME=/home/user"),
				map[string]string{
					"APP_DIR": "$XDG_CONFIG_HOME/app",
				},
			),
			inputs: []string{
				"$APP_DIR/.env",
				"$APP_DIR/env",
			},
			expected: []string{
				"/home/user/.config/app/.env",
				"/home/user/.config/app/env",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.environment.SubstitueAll(tc.inputs...)
			assert.EqualValues(t, actual, tc.expected)
		})
	}
}
