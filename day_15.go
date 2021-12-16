package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"
	"github.com/fatih/color"
	// "github.com/solarlune/paths"
)

var (
	//go:embed inputs/day_15.txt
	chitonFile string

	chitonWidth = 100

	chitonTest = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
)

func chiton() int {
	grid := parseChitonTiles(chitonFile)
	src := grid[0]
	target := grid[len(grid)-1]
	_, distance, found := astar.Path(target, src)
	if !found {
		fmt.Println("Failed to find path")
	}
	return int(distance)
}

func parseChitonTiles(f string) []*Tile {
	fullCaveWidth := chitonWidth
	fullCave := make([]*Tile, 0, fullCaveWidth*fullCaveWidth)

	lines := strings.Split(f, "\n")
	for row, line := range lines {
		for col, n := range line {
			nInt, _ := strconv.Atoi(string(n))
			t := &Tile{
				X:         col,
				Y:         row,
				Cost:      float64(nInt),
				FullWidth: fullCaveWidth,
			}
			fullCave = append(fullCave, t)
		}
	}
	for _, t := range fullCave {
		t.Grid = fullCave
	}
	return fullCave
}

// 2886 is too high, 2877 is too low
func chitonFullCave() int {
	grid := parseChitonTilesFull(chitonFile)

	src := grid[0]
	target := grid[len(grid)-1]

	_, distance, found := astar.Path(src, target)
	if !found {
		fmt.Println("Failed to find path")
	}
	fmt.Printf("Path risk is %d\n", int(distance))
	return 0
}

func parseChitonTilesFull(f string) []*Tile {
	fullCaveWidth := chitonWidth * 5
	fullCave := make([]*Tile, fullCaveWidth*fullCaveWidth)

	lines := strings.Split(f, "\n")
	for row, line := range lines {
		for col, n := range line {
			nInt, _ := strconv.Atoi(string(n))
			coord := fullCaveWidth*row + col
			t := &Tile{
				X:         col,
				Y:         row,
				Cost:      float64(nInt),
				FullWidth: fullCaveWidth,
			}
			fullCave[coord] = t

			for rowi := 0; rowi < 5; rowi++ {
				for coli := 0; coli < 5; coli++ {
					coord := fullCaveWidth*(row+(chitonWidth*rowi)) + (col + (chitonWidth * coli))
					nextVal := ((nInt + rowi + coli) % 9)
					if nextVal == 0 {
						nextVal = 9
					}
					t := &Tile{
						X:         col + (chitonWidth * coli),
						Y:         row + (chitonWidth * rowi),
						Cost:      float64(nextVal),
						FullWidth: fullCaveWidth,
					}
					fullCave[coord] = t
				}
			}
		}
	}

	for _, t := range fullCave {
		t.Grid = fullCave
	}
	return fullCave
}

func printChiton(grid []*Tile, width int) {
	for i, t := range grid {
		if t.Path {
			fmt.Print(color.BlueString("%d", int(t.Cost)))
		} else {
			fmt.Print(color.HiWhiteString("%d", int(t.Cost)))
		}
		if (i+1)%width == 0 {
			fmt.Println()
		}
	}
}

type Tile struct {
	X, Y      int
	Cost      float64
	Grid      []*Tile
	Path      bool
	FullWidth int
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0, 4)
	if up := t.Up(); up != nil {
		neighbors = append(neighbors, up)
	}
	if right := t.Right(); right != nil {
		neighbors = append(neighbors, right)
	}
	if left := t.Left(); left != nil {
		neighbors = append(neighbors, left)
	}
	if down := t.Down(); down != nil {
		neighbors = append(neighbors, down)
	}
	return neighbors
}

func (t *Tile) Up() *Tile {
	coord := t.FullWidth*(t.Y-1) + t.X
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Right() *Tile {
	if (t.X + 1) >= t.FullWidth {
		return nil
	}

	coord := t.FullWidth*t.Y + t.X + 1
	if (coord / t.FullWidth) != ((coord - 1) / t.FullWidth) {
		return nil
	}

	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Left() *Tile {
	if t.X == 0 {
		return nil
	}
	coord := t.FullWidth*t.Y + t.X - 1
	if (coord / t.FullWidth) != ((coord + 1) / t.FullWidth) {
		return nil
	}
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Down() *Tile {
	coord := t.FullWidth*(t.Y+1) + t.X
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return toT.Cost
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := math.Abs(float64(toT.X - t.X))
	absY := math.Abs(float64(toT.Y - t.Y))
	return absX + absY
}
