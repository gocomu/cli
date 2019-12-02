package cli

// RTout type helps cli's -out flag
type RTout int

const (
	// PortAudio output
	PortAudio RTout = iota
	// Oto output
	Oto
)

// ProjectType type helps cli's `cli` subcommand
type ProjectType int

const (
	// Cli project type
	Cli = ProjectType(iota)
	// Gui project type
	Gui
)
