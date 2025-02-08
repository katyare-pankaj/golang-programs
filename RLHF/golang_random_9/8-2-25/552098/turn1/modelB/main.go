package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	numDownloads  = 5
	downloadSpeed = 100 * 1024 * 1024 // 100 MB/s
)

type Download struct {
	url        string
	fileName   string
	downloader *Downloader
}

func (d *Download) Download() error {
	resp, err := http.Get(d.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(d.fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d *Download) Cleanup() {
	err := os.Remove(d.fileName)
	if err != nil {
		log.Printf("Error cleaning up %s: %v\n", d.fileName, err)
	} else {
		log.Printf("Cleaned up %s\n", d.fileName)
	}
}

type Downloader struct {
	downloadChan chan *Download
	wg           *sync.WaitGroup
}

func NewDownloader() *Downloader {
	downloadChan := make(chan *Download, numDownloads)
	wg := &sync.WaitGroup{}

	d := &Downloader{
		downloadChan: downloadChan,
		wg:           wg,
	}

	wg.Add(numDownloads)
	for i := 0; i < numDownloads; i++ {
		go d.downloadWorker()
	}

	return d
}

func (d *Downloader) AddDownload(download *Download) {
	d.downloadChan <- download
}

func (d *Downloader) Wait() {
	d.wg.Wait()
	close(d.downloadChan)
}

func (d *Downloader) downloadWorker() {
	defer d.wg.Done()

	for download := range d.downloadChan {
		err := download.Download()
		if err != nil {
			log.Printf("Error downloading %s: %v\n", download.url, err)
			continue
		}

		log.Printf("Downloaded %s\n", download.fileName)
		// Simulate cleanup time
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		download.Cleanup()
	}
}

func main() {
	downloader := NewDownloader()

	for i := 1; i <= numDownloads; i++ {
		fileName := fmt.Sprintf("download%d.txt", i)
		url := fmt.Sprintf("https://example.com/large_file_%d.txt", i) // Replace this with actual URLs

		download := &Download{
			url:        url,
			fileName:   fileName,
			downloader: downloader,
		}

		downloader.AddDownload(download)
	}

	downloader.Wait()
	fmt.Println("All downloads completed and cleaned up.")
}
