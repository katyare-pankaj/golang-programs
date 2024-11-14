package main

import "fmt"
import "time"

const (
    m, n, p = 1000, 1000, 1000
)

func main() {
    var a, b, c [m][n]float64
    
    // Initialize matrices
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            a[i][j] = float64(i*j + 1)
            b[i][j] = float64(j*i + 1)
        }
    }
    
    // Matrix multiplication
    start := time.Now()
    multiplyMatrices(a, b, &c)
    elapsed := time.Since(start)
    
    fmt.Printf("Time elapsed: %s\n", elapsed)
}

func multiplyMatrices(a, b [m][n]float64, c *[m][n]float64) {
    for i := 0; i < m; i++ {
        for j := 0; j < p; j++ {
            sum := 0.0
            for k := 0; k < n; k++ {
                sum += a[i][k] * b[k][j]
            }
            (*c)[i][j] = sum
        }
    }
}