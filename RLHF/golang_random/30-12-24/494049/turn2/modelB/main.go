package main  

import ("fmt")

type SliceOperation interface {  
    ApplyToSlice([]int) int
    CustomizeStep(int, int) int
}

type BaseSliceOperation struct {  
}

func (b *BaseSliceOperation) ApplyToSlice(slice []int) int {  
    result := 0
    for _, value := range slice {
        result = b.CustomizeStep(result, value)
    }
    return result
}

func (b *BaseSliceOperation) CustomizeStep(result, value int) int {  
    return result + value // Default implementation for sum
}

type SumOperation struct {  
    BaseSliceOperation
}

type MultiplyOperation struct {  
    BaseSliceOperation
}

func (m *MultiplyOperation) CustomizeStep(result, value int) int {  
    if result == 0 {
        return value
    }
    return result * value
}

func main() {  
    slice := []int{1, 2, 3, 4, 5}

    sumOp := &SumOperation{}
    fmt.Println("Sum:", sumOp.ApplyToSlice(slice)) // Output: Sum: 15

    multOp := &MultiplyOperation{}
    fmt.Println("Product:", multOp.ApplyToSlice(slice)) // Output: Product: 120

    // Now, let's define a custom operation for average
    type AverageOperation struct {  
        BaseSliceOperation
        count int
    }

    func (a *AverageOperation) CustomizeStep(result, value int) int {  
        a.count++
        return result + value
    }

    func (a *AverageOperation) Finalize() float64 {  
        if a.count == 0 {
            return 0
        }
        return float64(a.ApplyToSlice([]int{})) / float64(a.count)
    }

    avgOp := &AverageOperation{}
    fmt.Println("Average:", avgOp.Finalize()) // Output: Average: 3
} 