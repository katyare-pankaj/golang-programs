package main

import (
    "fmt"
    "time"
)

func producer(ch1 chan int, ch2 chan string) {
    for i := 1; i <= 10; i++ {
        select {
        case ch1 <- i:
            fmt.Println("Sent:", i, "to ch1")
        case ch2 <- fmt.Sprintf("%d", i):
            fmt.Println("Sent:", i, "to ch2")
        }
        time.Sleep(1 * time.Second)
    }
    close(ch1)
    close(ch2)
}

func consumer1(ch chan int) {
    for {
        if num, ok := <-ch; ok {
            fmt.Println("Received:", num, "from ch1")
        } else {
            break
        }
    }
}

func consumer2(ch chan string) {
    for {
        if str, ok := <-ch; ok {
            fmt.Println("Received:", str, "from ch2")