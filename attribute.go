package colors

// Styling attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FGBlack Attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

// Foreground Hi-Intensity text colors
const (
	FGHiBlack Attribute = iota + 90
	FGHiRed
	FGHiGreen
	FGHiYellow
	FGHiBlue
	FGHiMagenta
	FGHiCyan
	FGHiWhite
)

// Background text colors
const (
	BGBlack Attribute = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

// Background Hi-Intensity text colors
const (
	BGHiBlack Attribute = iota + 100
	BGHiRed
	BGHiGreen
	BGHiYellow
	BGHiBlue
	BGHiMagenta
	BGHiCyan
	BGHiWhite
)

// Attribute defines a single SGR Code
type Attribute int
