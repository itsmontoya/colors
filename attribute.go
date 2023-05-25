package colors

import (
	"fmt"
	"strconv"
)

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

func (a Attribute) String() string {
	return strconv.Itoa(int(a))
}

// Sprint will return a colored version of the value
func (a Attribute) Sprint(value string) (out string) {
	return fmt.Sprintf(colorFormat, a.String(), value, Reset)
}

// Sprintf will return a colored version for the format and values
func (a Attribute) Sprintf(format string, values ...interface{}) (out string) {
	value := fmt.Sprintf(format, values...)
	return a.Sprint(value)
}

// Print will output colored version of a given string
func (a Attribute) Print(value string) {
	out := a.Sprint(value)
	fmt.Print(out)
}

// Printf will output colored version of a given format and values
func (a Attribute) Printf(format string, values ...interface{}) {
	out := a.Sprintf(format, values...)
	fmt.Print(out)
}

// Printf will output colored version of a given format and values
func (a Attribute) Println(format string, values ...interface{}) {
	out := a.Sprintf(format, values...)
	fmt.Println(out)
}
