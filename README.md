# Colors
Colors is a helper library to add color to text

## Usage 
### Configuration.Sprint
```go
func ExampleConfiguration_Sprint() {
	c := New(Bold, FGGreen)
	out := c.Sprint("Hello world!\n")
	fmt.Printf("Sprint: %s", out)
}
```

### Configuration.Sprintf
```go
func ExampleConfiguration_Sprintf() {
	c := New(Bold, FGGreen)
	out := c.Sprintf("Hello world! My name is %s %s\n", "John", "Doe")
	fmt.Printf("Sprintf: %s", out)
}
```

### Configuration.Print
```go
func ExampleConfiguration_Print() {
	c := New(Bold, FGGreen)
	c.Print("Hello world!\n")
}
```

### Configuration.Printf
```go
func ExampleConfiguration_Printf() {
	c := New(Bold, FGGreen)
	c.Printf("Hello world! My name is %s %s\n", "John", "Doe")
}
```