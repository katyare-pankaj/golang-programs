package main  
import (  
    "fmt"
    "strings"
    "sync"
)

// Tokenize takes a string and returns a slice of tokens (words)
func Tokenize(text string) []string {
    return strings.Fields(text)
}

// nGramGenerator is a higher-order function that generates N-grams using a sliding window approach.
// It uses lazy evaluation by generating N-grams only when needed.
func nGramGenerator(tokens []string, n int) func() [][]string {
    return func() [][]string {
        result := make([][]string, 0)
        if n <= 0 {
            return result
        }
        for i := 0; i < len(tokens) - n + 1; i++ {
            result = append(result, tokens[i:i+n])
        }
        return result
    }
}

// parallelNGrams processes N-gram generation in parallel using a goroutine pool.
func parallelNGrams(tokens []string, n int, numWorkers int) [][]string {
    // Create a channel to receive generated N-grams
    ngramChan := make(chan [][]string, numWorkers)
    
    // Start worker goroutines
    wg := &sync.WaitGroup{}
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                select {
                case ngrams := <-ngramChan:
                    // Receive N-grams from the channel and append to the result
                case <-done:
                    return
                }
            }
        }()
    }
    
    // Distribute N-gram generation tasks among worker goroutines
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(ngramChan)
        gramGen := nGramGenerator(tokens, n)
        for _ = range gramGen() {
            ngramChan <- gramGen()
        }
    }()
    
    wg.Wait()
    return []][]string{} // Return an empty slice since all N-grams were processed via the channel
}

func main() {
    text := "Hello, world. This is a test sentence. Hello world again."
    tokens := Tokenize(text)
    
    const n = 2
    const numWorkers = 4 // Number of goroutine workers for parallelization
    
    //Using Partial function
    parallelNGramsFunc := func(tokens []string) [][]string {
        return parallelNGrams(tokens, n, numWorkers)
    }
    
    // Generate N-grams in parallel
    fmt.Println("Parallel Bigrams: ", parallelNGramsFunc(tokens))
} 