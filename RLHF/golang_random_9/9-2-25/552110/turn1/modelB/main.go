package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// MusicPlayer represents the state of the music player
type MusicPlayer struct {
	Playlist  []string
	Current   int
	Playing   bool
	SkipCount int
}

// NewMusicPlayer creates a new MusicPlayer with an empty playlist
func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{
		Playlist:  []string{},
		Current:   -1,
		Playing:   false,
		SkipCount: 0,
	}
}

// Play starts playing the current song in the playlist
func (p *MusicPlayer) Play() {
	if p.Current < 0 || p.Current >= len(p.Playlist) {
		fmt.Println("Playlist is empty.")
		return
	}

	p.Playing = true
	fmt.Printf("Playing: %s\n", p.Playlist[p.Current])

	// Simulate song playing for 2 seconds
	time.Sleep(2 * time.Second)
	p.Playing = false
	fmt.Println("Song finished.")
}

// Pause pauses the currently playing song
func (p *MusicPlayer) Pause() {
	if p.Playing {
		p.Playing = false
		fmt.Println("Song paused.")
	} else {
		fmt.Println("No song is playing.")
	}
}

// Skip skips to the next song in the playlist
func (p *MusicPlayer) Skip() {
	if p.Current < 0 || p.Current >= len(p.Playlist)-1 {
		fmt.Println("Playlist is empty or at the end.")
		return
	}

	p.Current++
	p.SkipCount++
	fmt.Printf("Skipping to: %s\n", p.Playlist[p.Current])
}

// AddSong adds a song to the playlist
func (p *MusicPlayer) AddSong(song string) {
	p.Playlist = append(p.Playlist, song)
	if p.Current < 0 {
		p.Current = 0
	}
	fmt.Printf("Added song: %s\n", song)
}

func main() {
	player := NewMusicPlayer()

	// Add some songs to the playlist
	player.AddSong("Song 1")
	player.AddSong("Song 2")
	player.AddSong("Song 3")

	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("Exiting music player...")
				os.Exit(0)
			case syscall.SIGUSR1: // Play/Pause
				if player.Playing {
					player.Pause()
				} else {
					player.Play()
				}
			case syscall.SIGUSR2: // Skip
				player.Skip()
			}
		}
	}()

	fmt.Println("Music player started. Use SIGUSR1 to play/pause, SIGUSR2 to skip, and SIGINT/SIGTERM to exit.")

	// Loop indefinitely to handle signals
	for {
		time.Sleep(1 * time.Second)
	}
}
