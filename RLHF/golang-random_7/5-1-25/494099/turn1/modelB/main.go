func sprintfExample() string {
	// Using fmt.Sprintf
	return fmt.Sprintf("Hello, %s! Your age is %d.", "Alice", 25)
}

func stringManipulationExample() string {
	// String manipulation using + operator
	name := "Bob"
	age := 30
	return "Hello, " + name + "! Your age is " + strconv.Itoa(age) + "."
}

func main() {
	fmt.Println(sprintfExample())            // Output: Hello, Alice! Your age is 25.
	fmt.Println(stringManipulationExample()) // Output: Hello, Bob! Your age is 30.
}