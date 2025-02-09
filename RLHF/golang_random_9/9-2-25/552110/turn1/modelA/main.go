package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type MusicPlayer struct {
	playlist []string
	current  int
	playing  bool
	mu       sync.Mutex
}

func NewMusicPlayer(playlist []string) *MusicPlayer {
	return &MusicPlayer{
		playlist: playlist,
		current:  0,
		playing:  false,
	}
}

func (mp *MusicPlayer) Play() error {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if mp.playing {
		return errors.New("already playing")
	}

	if len(mp.playlist) == 0 {
		return errors.New("playlist is empty")
	}

	mp.playing = true
	fmt.Printf("Playing: %s\n", mp.playlist[mp.current])
	// Simulate playing the current track
	go func() {
		time.Sleep(3 * time.Second)
		mp.Skip()
	}()

	return nil
}

func (mp *MusicPlayer) Pause() error {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if !mp.playing {
		return errors.New("not playing")
	}

	mp.playing = false
	fmt.Println("Paused")
	return nil
}

func (mp *MusicPlayer) Skip() error {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if len(mp.playlist) == 0 {
		return errors.New("playlist is empty")
	}

	mp.current = (mp.current + 1) % len(mp.playlist)
	fmt.Printf("Skipped to: %s\n", mp.playlist[mp.current])

	if mp.playing {
		// Simulate the next track playing
		go func() {
			time.Sleep(3 * time.Second)
			mp.Skip()
		}()
	}

	return nil
}

func main() {
	// Example playlist
	playlist := []string{"Song 1", "Song 2", "Song 3"}

	player := NewMusicPlayer(playlist)

	// Channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)

	fmt.Println("Music Player started. Use signals to control the player:")
	fmt.Println("SIGUSR1: Play/Pause toggle")
	fmt.Println("SIGUSR2: Skip")
	fmt.Println("SIGTERM: Exit")

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGUSR1:
			if player.playing {
				player.Pause()
			} else {
				if err := player.Play(); err != nil {
					fmt.Println("Error:", err)
				}
			}
		case syscall.SIGUSR2:
			if err := player.Skip(); err != nil {
				fmt.Println("Error:", err)
			}
		case syscall.SIGTERM:
			fmt.Println("Exiting...")
			return
		}
	}
}
