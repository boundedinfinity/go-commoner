package sheller

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

func Detect() Shell {
	found := Shells.Unknown

	switch runtime.GOOS {
	case "windows":
		found = detectWindowsShell()
	default:
		found = detectUnixShell()
	}

	return found
}

func DetectOk() (Shell, bool) {
	found := Detect()
	return found, found != Shells.Unknown
}

var (
	ErrShell   = errorer.New("shell error")
	errShellFn = errorer.Func(ErrShell)
)

func DetectErr() (Shell, error) {
	found := Detect()

	if found == Shells.Unknown {
		return found, errShellFn("shell not detected")
	}

	return found, nil
}

func detectUnixShell() Shell {
	if os.Getenv("ZSH_VERSION") != "" {
		return Shells.Zsh
	}

	if os.Getenv("BASH_VERSION") != "" {
		return Shells.Bash
	}

	if os.Getenv("FISH_VERSION") != "" {
		return Shells.Fish
	}

	if shell := os.Getenv("SHELL"); shell != "" {
		shell = filepath.Base(shell)
		switch shell {
		case string(Shells.Bash), string(Shells.Sh):
			return Shells.Bash
		case string(Shells.Zsh):
			return Shells.Zsh
		case string(Shells.Fish):
			return Shells.Fish
		case string(Shells.Csh):
			return Shells.Csh
		case string(Shells.Tcsh):
			return Shells.Tcsh
		}
	}

	if len(os.Args) > 0 {
		switch filepath.Base(os.Args[0]) {
		case string(Shells.Bash), string(Shells.Sh):
			return Shells.Bash
		case string(Shells.Zsh):
			return Shells.Zsh
		case string(Shells.Fish):
			return Shells.Fish
		case string(Shells.Csh):
			return Shells.Csh
		case string(Shells.Tcsh):
			return Shells.Tcsh
		}
	}

	return Shells.Unknown
}

func detectWindowsShell() Shell {
	comspec := strings.ToLower(os.Getenv("ComSpec"))

	if os.Getenv("PSModulePath") != "" || strings.Contains(comspec, "pwsh") {
		return Shells.Powershell
	}

	if strings.Contains(comspec, "cmd.exe") {
		return Shells.Cmd
	}

	if os.Getenv("MSYSTEM") != "" {
		return Shells.Bash
	}

	if os.Getenv("WSL_DISTRO_NAME") != "" {
		return detectUnixShell()
	}

	return Shells.Unknown
}
