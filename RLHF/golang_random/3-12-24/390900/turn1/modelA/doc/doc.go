package doc

// +godoc
import (
	"fmt"
	"go-programs/RLHF/golang_random/3-12-24/390900/turn1/modelA/utils"
)

func ExampleDouble() {
	result := utils.Double(10)
	fmt.Println(result) // Output: 20
}

func ExampleAdd() {
	result := utils.Add(10, 20)
	fmt.Println(result) // Output: 30
}
