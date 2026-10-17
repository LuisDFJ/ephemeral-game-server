package ansi

type Ansi struct {
	code string
}

func New(code string) Ansi {
	return Ansi{code:code}
}

func (a *Ansi) Print(s string) string {
	return "\x1b[" + a.code + s + "\x1b[0m"
}

// Bright white foreground Bright Black background
var LIGHT Ansi = New("97;100m")
var ERROR Ansi = New("91m")
var INFO Ansi = New("92m")
var LOADER Ansi = New("96m")
