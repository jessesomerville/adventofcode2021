package week3

import (
	"fmt"
	"strings"
)

// ImageEnhancer provides functions for enhancing the input image.
type ImageEnhancer struct {
	Image     []int
	Algorithm []int
	Width     int
}

// Enhance applies the enhancement algorithm to the input image.
func (ie *ImageEnhancer) Enhance() {
	// ie.Scale()
	enhanced := make([]int, 0, len(ie.Image))
	for i := range ie.Image {
		enhanced = append(enhanced, ie.NextPixel(i))
	}
	ie.Image = enhanced
}

// NextPixel takes and index of a pixel in the input image and returns the 9-bit binary number
// that can be used to lookup the new pixel value in the enhancement algorithm.
func (ie *ImageEnhancer) NextPixel(idx int) int {
	rowNum := idx / ie.Width
	colNum := idx % ie.Width
	if rowNum == 0 || rowNum == ie.Width-1 || colNum == 0 || colNum == ie.Width-1 {
		return 0
	}

	var lookup int
	tenths := 8
	for row := rowNum - 1; row < rowNum+2; row++ {
		for col := colNum - 1; col < colNum+2; col++ {
			coord := ie.Width*row + col
			if ie.Image[coord] == 1 {
				lookup += 1 << tenths
			}
			tenths--
		}
	}
	return ie.Algorithm[lookup]
}

// PixelCount returns the number of pixels that are 'on' in the image.
func (ie *ImageEnhancer) PixelCount() int {
	var cnt int
	for _, pixel := range ie.Image {
		cnt += pixel
	}
	return cnt
}

// Scale adds padding to the input image to provide space for enhancement.
func (ie *ImageEnhancer) Scale(amount int) {
	newWidth := ie.Width + (amount * 2)
	scaled := make([]int, 0, newWidth*newWidth)

	for rowNum := 0; rowNum < newWidth; rowNum++ {
		for colNum := 0; colNum < newWidth; colNum++ {
			if rowNum < amount || rowNum > newWidth-(amount+1) || colNum < amount || colNum > newWidth-(amount+1) {
				scaled = append(scaled, 0)
			} else {
				coord := ie.Width*(rowNum-amount) + (colNum - amount)
				scaled = append(scaled, ie.Image[coord])
			}
		}
	}
	ie.Image = scaled
	ie.Width = newWidth
}

// Trim removes the superfluous padding left after enhancement.
func (ie *ImageEnhancer) Trim(amount int) {
	newWidth := ie.Width - amount*2
	trimmed := make([]int, 0, newWidth*newWidth)
	for rowNum := 0; rowNum < newWidth; rowNum++ {
		for colNum := 0; colNum < newWidth; colNum++ {
			coord := ie.Width*(rowNum+amount) + (colNum + amount)
			trimmed = append(trimmed, ie.Image[coord])
		}
	}
	ie.Image = trimmed
	ie.Width = newWidth
}

// NewImageEnhancer returns a new ImageEnhancer for the given puzzle input.
func NewImageEnhancer(f string) *ImageEnhancer {
	parts := strings.Split(f, "\n\n")
	imageRows := strings.Split(parts[1], "\n")
	algo := make([]int, 0, 512)
	for _, r := range parts[0] {
		if r == '#' {
			algo = append(algo, 1)
		} else {
			algo = append(algo, 0)
		}
	}

	ie := &ImageEnhancer{
		Image:     make([]int, 0, len(parts[1])),
		Algorithm: algo,
		Width:     len(imageRows[0]),
	}

	for _, row := range imageRows {
		for _, pixel := range row {
			var p int
			if pixel == '#' {
				p = 1
			}

			ie.Image = append(ie.Image, p)
		}
	}
	return ie
}

func (ie *ImageEnhancer) String() string {
	sm := map[int]string{
		0: ".",
		1: "#",
	}
	buf := new(strings.Builder)
	for i, pixel := range ie.Image {
		fmt.Fprint(buf, sm[pixel])
		if (i+1)%ie.Width == 0 {
			fmt.Fprintln(buf)
		}
	}
	fmt.Fprintln(buf)
	return buf.String()
}
