package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

func main() {
	// Open the music file
	f, err := os.Open("example.mp3")
	if err != nil {
		log.Fatalf("failed to open music file: %v", err)
	}
	defer f.Close()

	// Create a new MP3 stream decoder
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalf("failed to decode music file: %v", err)
	}
	defer streamer.Close()

	// Create a new speaker
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Play the music file
	ctrl := beep.Play(streamer)

	// Wait for interrupt signal (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Stop playing the music file
	ctrl.Close()
	fmt.Println("Music player stopped.")
}
