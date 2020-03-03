package colors

import (
	"fmt"
	"testing"
)

func TestConfiguration_Sprint(t *testing.T) {
	c := New(Bold, FGGreen)
	out := c.Sprint("Hello world!\n")
	fmt.Printf("Sprint: %s", out)
}

func TestConfiguration_Sprintf(t *testing.T) {
	c := New(Bold, FGGreen)
	out := c.Sprintf("Hello world! My name is %s %s\n", "John", "Doe")
	fmt.Printf("Sprintf: %s", out)
}

func TestConfiguration_Print(t *testing.T) {
	c := New(Bold, FGGreen)
	c.Print("Hello world!\n")
}

func TestConfiguration_Printf(t *testing.T) {
	c := New(Bold, FGGreen)
	c.Printf("Hello world! My name is %s %s\n", "John", "Doe")
}
