package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// MusicPlayer represents the state of the music player
type MusicPlayer struct {
	Playlist []string
	Current  int
	Playing  bool
}

// NewMusicPlayer creates a new MusicPlayer with an empty playlist
func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{
		Playlist: []string{},
		Current:  -1,
		Playing:  false,
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

	fmt.Println("Music player started. Type 'play', 'pause', 'skip', or 'exit' to control the player.")

	// Read from standard input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		command := strings.TrimSpace(input)

		switch command {
		case "play":
			if !player.Playing {
				player.Play()
			} else {
				fmt.Println("Song is already playing.")
			}
		case "pause":
			player.Pause()
		case "skip":
			player.Skip()
		case "exit":
			fmt.Println("Exiting music player...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Please use 'play', 'pause', 'skip', or 'exit'.")
		}
	}
}
