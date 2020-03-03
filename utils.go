package colors

import (
	"strconv"
	"strings"
)

// makeSequence returns a formatted SGR sequence to be plugged into a color format string
func makeSequence(as []Attribute) string {
	format := make([]string, len(as))
	for i, v := range as {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}
