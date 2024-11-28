package calculator

// Public struct to represent a calculator
type Calculator struct {
	memory float64
}

// Public function to create a new calculator
func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

// Public function to add a number to the calculator's memory
func (c *Calculator) Add(num float64) {
	c.memory += num
}

// Public function to subtract a number from the calculator's memory
func (c *Calculator) Subtract(num float64) {
	c.memory -= num
}

// Public function to multiply the calculator's memory by a number
func (c *Calculator) Multiply(num float64) {
	c.memory *= num
}

// Public function to divide the calculator's memory by a number
func (c *Calculator) Divide(num float64) {
	if num == 0 {
		panic("division by zero")
	}
	c.memory /= num
}

// Public function to get the value in the calculator's memory
func (c *Calculator) GetMemory() float64 {
	return c.memory
}

// Private function to clear the calculator's memory
func (c *Calculator) clearMemory() {
	c.memory = 0
}
