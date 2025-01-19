package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	directory := "./data"           // Replace with your target directory
	ageLimit := 30 * 24 * time.Hour // 30 days
	deleteFilesOlderThan(directory, ageLimit)
	log.Println("Done")
}

func deleteFilesOlderThan(dir string, ageLimit time.Duration) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() {
			path := filepath.Join(dir, f.Name())
			info, err := os.Stat(path)
			if err != nil {
				log.Printf("Error stat'ing file %s: %v", path, err)
				continue
			}
			if time.Since(info.ModTime()) > ageLimit {
				if err := os.Remove(path); err != nil {
					log.Printf("Error deleting file %s: %v", path, err)
				} else {
					fmt.Printf("Deleted: %s\n", path)
				}
			}
		}
	}
}
