package main

import (
    "fmt"
    "sync"
    "time"
)

type Event struct {
    Data string
}

func processEvent(event Event, workChannel chan<- Event) {
    defer workChannel<-event
    fmt.Printf("Event received: %s\n", event.Data)
    time.Sleep(1 * time.Second) // Simulate work
}

func eventProcessorPool(wg *sync.WaitGroup, workChannel chan Event, numHandlers int) {
    defer wg.Done()
    for {
        event, ok := <-workChannel
        if !ok {
            return
        }
        processEvent(event, workChannel)
    }
}

func main() {
    wg := &sync.WaitGroup{}
    workChannel := make(chan Event, 100)
    numHandlers := 5

    wg.Add(numHandlers)
    for i := 0; i < numHandlers; i++ {
        go eventProcessorPool(wg, workChannel, numHandlers)
    }

    events := []Event{
        {Data: "Event 1"},
        {Data: "Event 2"},
        {Data: "Event 3"},
        {Data: "Event 4"},
    }

    for _, event := range events {
        processEvent(event, workChannel)
    }

    close(workChannel)
    wg.Wait()
    fmt.Println("All events processed.")
}
