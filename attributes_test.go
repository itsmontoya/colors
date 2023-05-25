package colors

import (
	"fmt"
	"testing"
)

func TestAttributes_Sprint(t *testing.T) {
	a := New(Bold, FGGreen)
	out := a.Sprint("Hello world!\n")
	fmt.Printf("Sprint: %s", out)
}

func TestAttributes_Sprintf(t *testing.T) {
	a := New(Bold, FGGreen)
	out := a.Sprintf("Hello world! My name is %s %s\n", "John", "Doe")
	fmt.Printf("Sprintf: %s", out)
}

func TestAttributes_Print(t *testing.T) {
	a := New(Bold, FGGreen)
	a.Print("Hello world!\n")
}

func TestAttributes_Printf(t *testing.T) {
	a := New(Bold, FGGreen)
	a.Printf("Hello world! My name is %s %s\n", "John", "Doe")
}

func ExampleAttributes_Sprint() {
	a := New(Bold, FGGreen)
	out := a.Sprint("Hello world!\n")
	fmt.Printf("Sprint: %s", out)
}

func ExampleAttributes_Sprintf() {
	a := New(Bold, FGGreen)
	out := a.Sprintf("Hello world! My name is %s %s\n", "John", "Doe")
	fmt.Printf("Sprintf: %s", out)
}

func ExampleAttributes_Print() {
	a := New(Bold, FGGreen)
	a.Print("Hello world!\n")
}

func ExampleAttributes_Printf() {
	a := New(Bold, FGGreen)
	a.Printf("Hello world! My name is %s %s\n", "John", "Doe")
}
