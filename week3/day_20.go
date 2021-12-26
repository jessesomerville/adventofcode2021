package week3

import (
	"fmt"
	"math"
	"strings"
)

// 5344 too high
func TrenchMap(f string) int {
	puzzleInput := parseImageEnhance(f)

	count := 0
	for _, pixel := range puzzleInput.InputImage {
		if pixel == '#' {
			count++
		}
	}
	fmt.Println(count)

	return 0
}

// ImageData represents the two pieces of the puzzle input.
type ImageData struct {
	EnhancementAlgorithm map[int]rune
	InputImage           []rune
	ImageWidth           int
}

func (id *ImageData) String() string {
	buf := new(strings.Builder)
	for i, bit := range id.InputImage {
		fmt.Fprint(buf, string(bit))
		if (i+1)%id.ImageWidth == 0 {
			fmt.Fprintln(buf)
		}
	}
	fmt.Fprintln(buf)
	return buf.String()
}

func (id *ImageData) Enhance() {
	id.Scale()
	outputImage := make([]rune, 0, len(id.InputImage))

	for row := 0; row < id.ImageWidth; row++ {
		for col := 0; col < id.ImageWidth; col++ {
			if row == 0 || row == id.ImageWidth-1 {
				outputImage = append(outputImage, '.')
				continue
			}
			if col == 0 || col == id.ImageWidth-1 {
				outputImage = append(outputImage, '.')
				continue
			}
			encodedPos := get3x3Int(col, row, id.InputImage)
			outputImage = append(outputImage, id.EnhancementAlgorithm[encodedPos])
		}
	}
	id.InputImage = outputImage
}

func get3x3Int(x, y int, in []rune) int {
	minY, maxY := y-1, y+1
	minX, maxX := x-1, x+1
	gridWidth := int(math.Sqrt(float64(len(in))))

	pos := 8
	num := 0
	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			coord := gridWidth*row + col
			if in[coord] == '#' {
				num += 1 << pos
			}
			pos--
		}
	}
	return num
}

func (id *ImageData) Scale() {
	newImageWidth := id.ImageWidth + 4
	scaled := make([]rune, newImageWidth*newImageWidth)

	for i := 0; i < len(scaled); i++ {
		row := i / newImageWidth
		col := i % newImageWidth

		mappedRow := row - 2
		mappedCol := col - 2

		if mappedRow >= 0 && mappedRow < id.ImageWidth && mappedCol >= 0 && mappedCol < id.ImageWidth {
			linear := id.ImageWidth*mappedRow + mappedCol
			scaled[i] = id.InputImage[linear]
		} else {
			scaled[i] = '.'
		}
	}
	id.InputImage = scaled
	id.ImageWidth = newImageWidth
}

func parseImageEnhance(f string) *ImageData {
	parts := strings.Split(f, "\n\n")

	enAlgStr := parts[0]
	enAlgMap := make(map[int]rune, 512)

	for i, r := range enAlgStr {
		enAlgMap[i] = r
	}

	inImage := []rune{}
	imageRows := strings.Split(parts[1], "\n")
	for _, row := range imageRows {
		for _, col := range row {
			inImage = append(inImage, col)
		}
	}

	return &ImageData{
		EnhancementAlgorithm: enAlgMap,
		InputImage:           inImage,
		ImageWidth:           len(imageRows),
	}
}
