
package main  
import (  
    "fmt"
    "strings"
    "sync"
    "time"
)
func wordCount(input string, wordCountMap map[string]int) {
    words := strings.Fields(input)
    for _, word := range words {
        word = strings.Trim(word, ",.;:") // Remove punctuation
        wordCountMap[word]++
    }
}
func main() {
    // Initialize the input string
    input := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor quam id massa faucibus dignissim. Nam viverra euismod odio, gravida pellentesque urna varius vitae. Sed auctor ultrices ex, vel lacinia dolor. Vestibulum id ligula porta felis euismod semper. Nullam eget metus id nulla ultricies commodo. Proin sed tortor molestie, maximus ligula id, elementum metus.
    Donec ullamcorper nulla non metus auctor fringilla. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Nullam id dolor id nibh ultricies vehicula ut id elit. Aenean eu leo quam. Pellentesque ornare sem lacinia quam venenatis vestibulum. Nam sagittis dictum nisl, sed ullamcorper ipsum dignissim ac.
    Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Vestibulum luctus, nisl vitae scelerisque malesuada, metus odio varius elit, a euismod diam lacus vestibulum turpis. Nullam quis risus eget urna mollis ornare vel eu leo.`
    
    // Create a map to store word counts
    wordCountMap := make(map[string]int)

    numGoRoutines := 10  // Set the number of go routines

    // Start profiling
    pprofStart()

    start := time.Now()
    var wg sync.WaitGroup
    wg.Add(numGoRoutines)

    // Distribute the word counting task across multiple go routines
    for i := 0; i < numGoRoutines; i++ {
        go func(id int) {
            wordCount(input, wordCountMap)
            wg.Done()
        }(i)
    }

    wg.Wait()
    end := time.Now()
    duration := end.Sub(start)
    fmt.Printf("Total time taken: %s\n", duration)
    // Stop profiling
    pprofStop()

    // Display the top 10 hot functions after profiling
    pprofTop10()
}
// pprofStart starts CPU profiling.
func pprofStart() {
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
}
// pprofStop stops CPU profiling and writes the profile data to a file.
func pprofStop() {
    pprof.StopCPUProfile()
    f.Close()
}
// pprofTop10 displays the top 10 hot functions based on the CPU profile.
func pprofTop10() {
    f, err := os.Open("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    p, err := pprof.Parse(f)
    if err != nil {
        log.Fatal(err)
    }

    top10 := p.Top10()
    fmt.Println("Top 10 Hot Functions:")
    for _, fn := range top10 {
        fmt.Printf("%s: %d\n", fn.Name, fn.Count)