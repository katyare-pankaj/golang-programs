package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// LazyImage represents an image that will be computed lazily.
type LazyImage struct {
	width, height int
	compute       func() *image.RGBA
	image         *image.RGBA
	done          bool
}

// NewLazyImage creates a new LazyImage with the given dimensions and computation function.
func NewLazyImage(width, height int, compute func() *image.RGBA) *LazyImage {
	return &LazyImage{
		width:   width,
		height:  height,
		compute: compute,
	}
}

// At returns the color of the pixel at the specified position in the image.
func (l *LazyImage) At(x, y int) color.Color {
	if !l.done {
		l.image = l.compute()
		l.done = true
	}
	return l.image.At(x, y)
}

// Bounds returns the image's bounds.
func (l *LazyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, l.width, l.height)
}

// ColorModel returns the image's color model.
func (l *LazyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// ApplyFilter applies a simple grayscale filter to the image lazily.
func (l *LazyImage) ApplyFilter() *LazyImage {
	return NewLazyImage(l.width, l.height, func() *image.RGBA {
		img := l.compute()
		filteredImg := image.NewRGBA(img.Bounds())
		for y := 0; y < img.Bounds().Max.Y; y++ {
			for x := 0; x < img.Bounds().Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
				filteredImg.Set(x, y, color.RGBA{gray, gray, gray, uint8(a)})
			}
		}
		return filteredImg
	})
}

func generateRandomImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(rand.Intn(256))
			g := uint8(rand.Intn(256))
			b := uint8(rand.Intn(256))
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return img
}

func main() {
	const imageWidth, imageHeight = 1000, 1000
	rand.Seed(time.Now().UnixNano())

	// Create a lazy image with a random noise generator
	originalImage := NewLazyImage(imageWidth, imageHeight, func() *image.RGBA {
		return generateRandomImage(imageWidth, imageHeight)
	})

	// Apply the grayscale filter lazily on the original image
	filteredImage := originalImage.ApplyFilter()

	// Now, if we need to access pixels from the filtered image, it will be computed lazily
	startTime := time.Now()
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			filteredImage.At(x, y) // Compute the filtered image pixels lazily
		}
	}
	fmt.Println("Time taken to process the image:", time.Since(startTime))
}
