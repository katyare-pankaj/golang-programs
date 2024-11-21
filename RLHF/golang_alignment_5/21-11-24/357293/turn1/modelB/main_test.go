package fibonacci
import "testing"
func Fibonacci(n int) int {
    // Base case: The first two Fibonacci numbers are 0 and 1.
    if n < 2 {
        return n
    }
    // Recursive case: Calculate the n-th Fibonacci number by calling the function with (n-1) and (n-2).
    return Fibonacci(n-1) + Fibonacci(n-2)
}
func TestFibonacci_BaseCase(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"First Fibonacci number", 0, 0},
        {"Second Fibonacci number", 1, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Fibonacci(tt.n); got != tt.want {
                t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
            }
        })
    }
}
func TestFibonacci_SmallNumbers(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"Fibonacci of 5", 5, 5},
        {"Fibonacci of 10", 10, 55},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Fibonacci(tt.n); got != tt.want {
                t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
            }
        })
    }
}
func TestFibonacci_LargeNumbers(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"Fibonacci of 35", 35, 9227465},
        {"Fibonacci of 40", 40, 102334155},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Fibonacci(tt.n); got != tt.want {