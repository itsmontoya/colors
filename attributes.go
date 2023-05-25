package colors

import "fmt"

// New will return a new
func New(as ...Attribute) Attributes {
	return as
}

type Attributes []Attribute

// Sprint will return a colored version of the value
func (a Attributes) Sprint(value string) (out string) {
	seq := makeSequence(a)
	return fmt.Sprintf(colorFormat, seq, value, Reset)
}

// Sprintf will return a colored version for the format and values
func (a Attributes) Sprintf(format string, values ...interface{}) (out string) {
	value := fmt.Sprintf(format, values...)
	return a.Sprint(value)
}

// Print will output colored version of a given string
func (a Attributes) Print(value string) {
	out := a.Sprint(value)
	fmt.Print(out)
}

// Printf will output colored version of a given format and values
func (a Attributes) Printf(format string, values ...interface{}) {
	out := a.Sprintf(format, values...)
	fmt.Print(out)
}

// Printf will output colored version of a given format and values
func (a Attributes) Println(format string, values ...interface{}) {
	out := a.Sprintf(format, values...)
	fmt.Println(out)
}
