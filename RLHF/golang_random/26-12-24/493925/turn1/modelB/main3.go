package main  
import (  
    "fmt"
    "math/rand"
    "time"
)
const (  
    numProducers   = 3
    numConsumers   = 2
    bufferSize     = 10
)
func producer(id int, data chan<- int, quit <-chan struct{}) {
    for {  
        select {  
        case <-quit:
            return // Gracefully quit if quit channel is closed
        default:
            value := rand.Intn(100)
            data <- value // Send data to the channel
            fmt.Printf("Producer %d sent %d\n", id, value)
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        }
    }
}

func consumer(id int, data <-chan int, quit <-chan struct{}) {
    for {  
        select {  
        case <-quit:
            return // Gracefully quit if quit channel is closed
        case value := <-data: // Receive data from the channel
            fmt.Printf("Consumer %d received %d\n", id, value)
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        }
    }
}

func main() {  
    data := make(chan int, bufferSize) // Buffered channel for producer-consumer pattern
    quit := make(chan struct{})