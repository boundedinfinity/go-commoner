package environmenter

import (
	"os"
	"strings"
)

func XdgDirs() map[string]string {
	return map[string]string{
		"XDG_CACHE_HOME":      "$HOME/.cache",
		"XDG_CONFIG_HOME":     "$HOME/.config",
		"XDG_DATA_HOME":       "$HOME/.local/share",
		"XDG_STATE_HOME":      "$HOME/.local/state",
		"XDG_RUNTIME_DIR":     "$HOME/.runtime",
		"XDG_DESKTOP_DIR":     "$HOME/Desktop",
		"XDG_DOWNLOAD_DIR":    "$HOME/Downloads",
		"XDG_TEMPLATES_DIR":   "$HOME/Templates",
		"XDG_PUBLICSHARE_DIR": "$HOME/Public",
		"XDG_DOCUMENTS_DIR":   "$HOME/Documents",
		"XDG_MUSIC_DIR":       "$HOME/Music",
		"XDG_PICTURES_DIR":    "$HOME/Pictures",
		"XDG_VIDEOS_DIR":      "$HOME/Videos",
	}
}

func CommonEnvFiles() func() map[string]string {
	return EnvFiles(".env", "$HOME/.env", "$XDG_CONFIG_HOME/.env")
}

func EnvFiles(paths ...string) func() map[string]string {
	return func() map[string]string {
		lines := []string{}

		for _, path := range paths {
			if bs, err := os.ReadFile(path); err == nil {
				lines = append(lines, strings.Split(string(bs), "\n")...)
			}
		}

		return KV(lines...)
	}
}

func EnvironmentVariable() map[string]string {
	return KV(os.Environ()...)
}

func KV(items ...string) map[string]string {
	mapping := make(map[string]string)

	for _, item := range items {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) == 2 {
			mapping[parts[0]] = parts[1]
		}
	}

	return mapping
}
