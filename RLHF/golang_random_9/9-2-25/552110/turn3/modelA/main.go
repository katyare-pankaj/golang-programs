package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var playlist = []string{"Song 1", "Song 2", "Song 3"} // Sample playlist
var currentTrack = 0
var playing = false

func main() {
	pipeName := `\\.\pipe\musicplayer`

	go runServer(pipeName)

	select {} // Keep the main function running
}

func runServer(pipeName string) {
	listener, err := net.Listen("unix", pipeName) // Use "unix" for cross-platform simulation; Windows requires more setup
	if err != nil {
		fmt.Println("Error creating named pipe:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Music player server started...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		handleCommand(command)
	}
}

func handleCommand(command string) {
	switch command {
	case "play":
		playing = true
		fmt.Printf("Playing: %s\n", playlist[currentTrack])
	case "pause":
		playing = false
		fmt.Println("Paused")
	case "skip":
		currentTrack = (currentTrack + 1) % len(playlist)
		fmt.Printf("Skipped to: %s\n", playlist[currentTrack])
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func playMusic() {
	for {
		if playing {
			fmt.Printf("Playing track: %s\n", playlist[currentTrack])
			time.Sleep(2 * time.Second) // Simulate playing
		}
	}
}
