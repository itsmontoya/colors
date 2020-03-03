package colors

import "fmt"

// New will return a new
func New(as ...Attribute) *Configuration {
	var c Configuration
	c.as = as
	return &c
}

// Configuration represnets a configuration of color and styling attributes
type Configuration struct {
	as []Attribute
}

// Sprint will return a colored version of the value
func (c *Configuration) Sprint(value string) (out string) {
	seq := makeSequence(c.as)
	return fmt.Sprintf(colorFormat, seq, value, Reset)
}

// Sprintf will return a colored version for the format and values
func (c *Configuration) Sprintf(format string, values ...interface{}) (out string) {
	value := fmt.Sprintf(format, values...)
	return c.Sprint(value)
}

// Print will output colored version of a given string
func (c *Configuration) Print(value string) {
	out := c.Sprint(value)
	fmt.Print(out)
}

// Printf will output colored version of a given format and values
func (c *Configuration) Printf(format string, values ...interface{}) {
	out := c.Sprintf(format, values...)
	fmt.Print(out)
}
