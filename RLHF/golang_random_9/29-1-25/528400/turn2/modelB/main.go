package main  
import (  
    "context"
    "fmt"
    "net/http"
    "os"
    "runtime"
    "runtime/pprof"
    "sync"
    "time"
)

// Task simulates a work task by sleeping for a short time
func Task(id int, ctx context.Context, workCh <-chan int) {
    for {
        select {
        case <-ctx.Done():
            // If context is canceled, return
            return
        case work := <-workCh:
            fmt.Printf("Task %d processing work: %d\n", id, work)
            // Simulate work taking time
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func main() {
    // Create a file to store the CPU profile
    f, err := os.Create("cpu.prof")
    if err != nil {
        fmt.Println("Could not create CPU profile:", err)
        return
    }
    defer f.Close()
    
    // Start CPU profiling
    if err := pprof.StartCPUProfile(f); err != nil {
        fmt.Println("Could not start CPU profile:", err)
        return
    }
    defer pprof.StopCPUProfile()

    numTasks := 10
    workCh := make(chan int, 100) // Buffered channel to reduce blocking
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Set GOMAXPROCS to the number of available CPUs
    runtime.GOMAXPROCS(runtime.NumCPU())

    // Start worker pool
    for i := 0; i < numTasks; i++ {
        go Task(i, ctx, workCh)
    }

    // Simulate producing work for tasks