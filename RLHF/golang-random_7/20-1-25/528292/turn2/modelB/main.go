package main  
import (  
    "fmt"
    "sort"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

type DataPoint struct {  
    Features []float64
    Label    string
}

func main() {  
    // Generate sample data points
    numPoints := 1000000
    data := make([]DataPoint, 0, numPoints)
    for i := 0; i < numPoints; i++ {  
        numFeatures := 10
        features := make([]float64, numFeatures)
        for j := 0; j < numFeatures; j++ {  
            features[j] = rand.Float64()
        }  
        label := fmt.Sprintf("Label_%d", i%3)
        data = append(data, DataPoint{Features: features, Label: label})
    }  
    
    // ------ Sort the data by Feature 3 for example ----------
    t := time.Now()
    sort.Slice(data, func(i, j int) bool {
	    return data[i].Features[3] < data[j].Features[3]
    })
    fmt.Println("Time taken to sort data:", time.Since(t))
    
    // ------ Example: Data processing using matrix operations ------
    // Create a feature matrix
    X := mat.NewDense(numPoints, numFeatures, nil)
    for i, dp := range data {  
        for j, feature := range dp.Features {  
            X.Set(i, j, feature)
        }  
    }  

    // ------ Example: Vectorized multiplication ------
    y := mat.NewDense(numPoints, 1, nil)
    w := mat.NewDense(numFeatures, 1, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0})
    t = time.Now()
    X.Mul(y, w)  // Use vectorized multiplication from gonum.org/v1/gonum/mat
    fmt.Println("Time taken for matrix multiplication:", time.Since(t))
    
    // ------ Search for specific label ------
    targetLabel := "Label_2"
    t = time.Now()
    start := sort.Search(len(data), func(i int) bool {
	    return data[i].Label >= targetLabel
    })