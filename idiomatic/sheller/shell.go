package sheller

type Shell string

func (this Shell) String() string {
	return string(this)
}

var Shells = shells{
	Unknown:    "unknown",
	Bash:       "bash",
	Cmd:        "cmd.exe",
	Csh:        "csh",
	Powershell: "powershell.exe",
	Fish:       "fish",
	Sh:         "sh",
	Tcsh:       "tcsh",
	Zsh:        "zsh",
}

type shells struct {
	Unknown    Shell
	Bash       Shell
	Cmd        Shell
	Csh        Shell
	Fish       Shell
	Powershell Shell
	Sh         Shell
	Tcsh       Shell
	Zsh        Shell
}
