package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

const pipeName = `\\.\pipe\musicPlayer`

func play() {
	fmt.Println("Playing music...")
	time.Sleep(5 * time.Second)
}

func pause() {
	fmt.Println("Pausing music...")
	time.Sleep(2 * time.Second)
}

func skip() {
	fmt.Println("Skipping music...")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "client" {
		clientMode()
	} else {
		serverMode()
	}
}

// Server Mode: Handles IPC and processes commands
func serverMode() {
	pipe, err := windows.CreateNamedPipe(
		syscall.StringToUTF16Ptr(pipeName),
		windows.PIPE_ACCESS_DUPLEX,
		windows.PIPE_TYPE_MESSAGE|windows.PIPE_READMODE_MESSAGE|windows.PIPE_WAIT,
		1, 1024, 1024, 0, nil,
	)

	if err != nil {
		log.Fatalf("Error creating pipe: %v", err)
	}
	defer windows.CloseHandle(pipe)

	fmt.Println("Waiting for client connection...")
	err = windows.ConnectNamedPipe(pipe, nil)
	if err != nil {
		log.Fatalf("Error connecting to pipe: %v", err)
	}

	fmt.Println("Client connected!")

	reader := bufio.NewReader(os.NewFile(uintptr(pipe), "pipe"))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading from pipe: %v", err)
		}
		signal := strings.TrimSpace(line)

		switch strings.ToLower(signal) {
		case "play":
			play()
		case "pause":
			pause()
		case "skip":
			skip()
		default:
			fmt.Println("Invalid command received:", signal)
		}
	}
}

// Client Mode: Sends a command to the server
func clientMode() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go client <play|pause|skip>")
		return
	}

	command := os.Args[2] + "\n"
	pipe, err := os.OpenFile(pipeName, os.O_WRONLY, 0)
	if err != nil {
		log.Fatalf("Error opening pipe: %v", err)
	}
	defer pipe.Close()

	_, err = pipe.WriteString(command)
	if err != nil {
		log.Fatalf("Error writing to pipe: %v", err)
	}

	fmt.Println("Command sent:", strings.TrimSpace(command))
}
