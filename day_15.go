package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"
	"github.com/fatih/color"
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

// 594 too high
func chiton() int {
	grid := parseChitonTiles(chitonFile)

	src := grid[0]
	target := grid[len(grid)-1]

	path, distance, found := astar.Path(src, target)
	if !found {
		fmt.Println("Failed to find path")
	}
	fmt.Printf("Path is %d tiles long\n", int(distance))

	for _, p := range path {
		tile := p.(*Tile)
		tile.Path = true
	}
	printChiton(grid)
	return 0
}

// func parseChitons(f string) []int {
// 	lines := strings.Split(f, "\n")
// 	grid := make([]int, len(lines)*len(lines[0]))

// 	for row, line := range lines {
// 		for col, r := range line {
// 			n, _ := strconv.Atoi(string(r))
// 			coord := len(lines[0])*row + col
// 			grid[coord] = n
// 		}
// 	}
// 	return grid
// }

func parseChitonTiles(f string) []*Tile {
	lines := strings.Split(f, "\n")
	grid := make([]*Tile, len(lines)*len(lines[0]))

	for row, line := range lines {
		for col, r := range line {
			n, _ := strconv.Atoi(string(r))
			t := &Tile{
				X:    col,
				Y:    row,
				Cost: float64(n),
			}
			coord := len(lines[0])*row + col
			grid[coord] = t
		}
	}

	for _, t := range grid {
		t.Grid = grid
	}
	return grid
}

func printChiton(grid []*Tile) {
	for i, t := range grid {
		if t.Path {
			fmt.Print(color.BlueString("%d", int(t.Cost)))
		} else {
			fmt.Print(color.HiWhiteString("%d", int(t.Cost)))
		}
		if (i+1)%chitonWidth == 0 {
			fmt.Println()
		}
	}
}

// func printChiton(grid []int) {
// 	for i, c := range grid {
// 		fmt.Print(c)
// 		if (i+1)%chitonWidth == 0 {
// 			fmt.Println()
// 		}
// 	}
// }

type Tile struct {
	X, Y int
	Cost float64
	Grid []*Tile
	Path bool
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
	coord := chitonWidth*(t.Y-1) + t.X
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Right() *Tile {
	if (t.X + 1) >= chitonWidth {
		return nil
	}
	coord := chitonWidth*t.Y + t.X + 1
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Left() *Tile {
	if t.X == 0 {
		return nil
	}
	coord := chitonWidth*t.Y + t.X + 1
	if coord < 0 || coord >= len(t.Grid) {
		return nil
	}
	return t.Grid[coord]
}

func (t *Tile) Down() *Tile {
	coord := chitonWidth*(t.Y+1) + t.X
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
