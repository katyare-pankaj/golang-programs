package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	pipeName = `\\.\pipe\musicPlayer`
)

func play() {
	// Simulate playing music
	fmt.Println("Playing music...")
	time.Sleep(5 * time.Second)
}

func pause() {
	// Simulate pausing music
	fmt.Println("Pausing music...")
	time.Sleep(2 * time.Second)
}

func skip() {
	// Simulate skipping music
	fmt.Println("Skipping music...")
}

func main() {
	go handleIPC()
	play()

	// Loop forever, pausing, skipping, or playing music based on the signals received
	for {
		select {
		case <-time.After(2 * time.Second):
			pause()
			play()
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func handleIPC() {
	// Create a named pipe for IPC
	pipe, err := os.Create(pipeName)
	if err != nil {
		log.Fatalf("Error creating pipe: %v", err)
	}
	defer pipe.Close()

	// Listen for signals from the pipe
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		signal := scanner.Text()
		switch strings.ToLower(signal) {
		case "play":
			play()
		case "pause":
			pause()
		case "skip":
			skip()
		default:
			fmt.Println("Invalid signal received:", signal)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from pipe: %v", err)
	}
}
