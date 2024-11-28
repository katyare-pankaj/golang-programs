// example.go
package main

func helloWorld() {
	// print the greeting
	println("Hello, World!")
}
func add(x, y int) (result int) {
	result = x + y
	return
}

var (
	numberOfTries = 10 // this is
	flavor        = "chocolate"
)

func processData(data []byte) int {
	byteCount := len(data)
	for i := 0; i < byteCount; i++ {
		data[i] = data[i] * 2
	}
	return byteCount
}
