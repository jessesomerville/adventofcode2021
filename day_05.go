package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day_05.txt
var vectorsFile string
var gridWidth = 1000

func getCoords(line string) (x1, y1, x2, y2 int) {
	elems := strings.Split(line, " ")
	p1, p2 := elems[0], elems[2]
	a := strings.Split(p1, ",")
	b := strings.Split(p2, ",")
	x1, _ = strconv.Atoi(a[0])
	x2, _ = strconv.Atoi(b[0])
	y1, _ = strconv.Atoi(a[1])
	y2, _ = strconv.Atoi(b[1])
	return x1, y1, x2, y2
}

func hydrothermalVenture() int {
	grid := make([]int, gridWidth*gridWidth)

	for line := range parseFile(vectorsFile) {
		x1, y1, x2, y2 := getCoords(line)
		if x1 == x2 {
			yMin, yMax := y1, y2
			if yMin > yMax {
				yMin, yMax = yMax, yMin
			}
			for i := yMin; i <= yMax; i++ {
				coord := gridWidth*i + x1
				grid[coord]++
			}
		} else if y1 == y2 {
			xMin, xMax := x1, x2
			if xMin > xMax {
				xMin, xMax = xMax, xMin
			}
			for i := xMin; i <= xMax; i++ {
				coord := gridWidth*y1 + i
				grid[coord]++
			}
		}
	}
	overlapping := 0
	for _, v := range grid {
		if v > 1 {
			overlapping++
		}
	}
	return overlapping
}

func hydrothermalVentureDiagonals() int {
	grid := make([]int, gridWidth*gridWidth)

	for line := range parseFile(vectorsFile) {
		x1, y1, x2, y2 := getCoords(line)
		yMin, yMax := y1, y2
		if yMin > yMax {
			yMin, yMax = yMax, yMin
		}
		xMin, xMax := x1, x2
		if xMin > xMax {
			xMin, xMax = xMax, xMin
		}
		if x1 == x2 {
			for i := yMin; i <= yMax; i++ {
				coord := gridWidth*i + x1
				grid[coord]++
			}
		} else if y1 == y2 {
			for i := xMin; i <= xMax; i++ {
				coord := gridWidth*y1 + i
				grid[coord]++
			}
		} else {
			slope := float64(y1-y2) / float64(x1-x2)
			if slope == 1 {
				for i := 0; i <= xMax-xMin; i++ {
					coord := gridWidth*(yMin+i) + (xMin + i)
					grid[coord]++
				}
			} else if slope == -1 {
				for i := 0; i <= xMax-xMin; i++ {
					coord := gridWidth*(yMax-i) + (xMin + i)
					grid[coord]++
				}
			}
		}
	}
	overlapping := 0
	for _, v := range grid {
		if v > 1 {
			overlapping++
		}
	}
	return overlapping
}
