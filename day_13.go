package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed inputs/day_13.txt
	paperFile string
)

func origami() {
	coords, folds, rowLen, colLen := parsePaper(paperFile)
	grid := make([]int, rowLen*colLen)
	for _, c := range coords {
		coord := rowLen*c.y + c.x
		grid[coord] = 1
	}

	for _, fold := range folds {
		if fold.x != 0 {
			grid = reflectX(grid, fold.x, rowLen)
			rowLen = fold.x
		} else {
			grid = reflectY(grid, fold.y, rowLen)
		}
	}
	printGrid(grid, rowLen)
}

func reflectY(grid []int, axis, rowLen int) []int {
	var topGrid, bottomGrid []int

	for i, c := range grid {
		if i/rowLen < axis {
			topGrid = append(topGrid, c)
		} else {
			bottomGrid = append(bottomGrid, c)
		}
	}

	rowCount := len(topGrid) / rowLen
	for i := 0; i < len(topGrid); i++ {
		row := i / rowLen
		btmRow := rowCount - row
		col := i % rowLen

		topCoord := rowLen*row + col
		btmCoord := rowLen*btmRow + col
		if btmCoord == len(bottomGrid) {
			break
		}
		if bottomGrid[btmCoord] == 1 {
			topGrid[topCoord] = 1
		}
	}

	return topGrid
}

func reflectX(grid []int, axis, rowLen int) []int {
	rowCount := len(grid) / rowLen

	leftGrid, rightGrid := make([]int, 0, axis*rowCount), make([]int, 0, axis*rowCount)
	for i, c := range grid {
		x := i % rowLen
		if x < axis {
			leftGrid = append(leftGrid, c)
		} else if x > axis {
			rightGrid = append(rightGrid, c)
		}
	}
	rowLen = axis

	for i := 0; i < len(rightGrid); i++ {
		row := i / rowLen
		col := i % rowLen
		leftCoord := rowLen*row + col
		rightCoord := rowLen*row + (rowLen - col - 1)
		if rightGrid[rightCoord] == 1 {
			leftGrid[leftCoord] = 1
		}
	}
	return leftGrid
}

type coord struct {
	x, y int
}

func parsePaper(f string) (coords, folds []*coord, maxX, maxY int) {
	parts := strings.Split(f, "\n\n")

	coords = make([]*coord, 0, len(parts[0]))
	folds = make([]*coord, 0, len(parts[1]))

	for _, c := range strings.Split(parts[0], "\n") {
		cStr := strings.Split(c, ",")
		coord := new(coord)
		coord.x, _ = strconv.Atoi(cStr[0])
		coord.y, _ = strconv.Atoi(cStr[1])
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
		coords = append(coords, coord)
	}

	for _, fs := range strings.Split(parts[1], "\n") {
		fStr := strings.Split(fs, "=")
		fold := new(coord)
		switch fStr[0] {
		case "x":
			fold.x, _ = strconv.Atoi(fStr[1])
		case "y":
			fold.y, _ = strconv.Atoi(fStr[1])
		}
		folds = append(folds, fold)
	}

	return coords, folds, maxX + 1, maxY + 1
}

func printGrid(grid []int, rowLen int) {
	for i, c := range grid {
		switch c {
		case 0:
			fmt.Print(".")
		case 1:
			fmt.Print("#")
		}
		if (i+1)%rowLen == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func printGridDiv(grid []int, rowLen, axis int, x bool) {
	for i, c := range grid {
		if x && i%rowLen == axis {
			fmt.Print("|")
		} else if !x && i/rowLen == axis {
			fmt.Print("-")
		} else {
			switch c {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			}
		}
		if (i+1)%rowLen == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
