package week4

import (
	"fmt"
	"strings"
)

// SeaCucumbers finds how many moves before no sea cucumbers can move.
func SeaCucumbers(f string) {
	scg := NewSCG(f)

	steps := 0
	for {
		steps++
		movedEast := scg.MigrateEast()
		movedSouth := scg.MigrateSouth()
		if !movedEast && !movedSouth {
			break
		}
	}
	fmt.Println(steps)
}

// SeaCucumberGrid represents the puzzle input.
type SeaCucumberGrid struct {
	Grid       [][]rune
	Rows, Cols int
}

// MigrateEast moves any able sea cucumbers East.
func (scg *SeaCucumberGrid) MigrateEast() bool {
	scgCopy := new(SeaCucumberGrid)
	scg.Copy(scgCopy)
	anyMoved := false
	for row := 0; row < scg.Rows; row++ {
		for col := 0; col < scg.Cols; col++ {
			spot := scg.Grid[row][col]
			nextSpot := scg.Grid[row][(col+1)%scg.Cols]
			if spot == '>' && nextSpot == '.' {
				scgCopy.Grid[row][col] = '.'
				scgCopy.Grid[row][(col+1)%scg.Cols] = '>'
				anyMoved = true
			}
		}
	}
	scgCopy.Copy(scg)
	return anyMoved
}

// MigrateSouth moves any able sea cucumbers South.
func (scg *SeaCucumberGrid) MigrateSouth() bool {
	scgCopy := new(SeaCucumberGrid)
	scg.Copy(scgCopy)
	anyMoved := false
	for row := 0; row < scg.Rows; row++ {
		for col := 0; col < scg.Cols; col++ {
			spot := scg.Grid[row][col]
			nextSpot := scg.Grid[(row+1)%scg.Rows][col]
			if spot == 'v' && nextSpot == '.' {
				scgCopy.Grid[row][col] = '.'
				scgCopy.Grid[(row+1)%scg.Rows][col] = 'v'
				anyMoved = true
			}
		}
	}
	scgCopy.Copy(scg)
	return anyMoved
}

// Copy makes a deep copy of the reciever.
func (scg SeaCucumberGrid) Copy(scgCopy *SeaCucumberGrid) {
	scgCopy.Grid = make([][]rune, 0, len(scg.Grid))
	scgCopy.Rows = scg.Rows
	scgCopy.Cols = scg.Cols
	for _, r := range scg.Grid {
		row := make([]rune, len(r))
		copy(row, r)
		scgCopy.Grid = append(scgCopy.Grid, row)
	}
}

// NewSCG creates a new SeaCucumberGrid from the puzzle input.
func NewSCG(f string) *SeaCucumberGrid {
	rows := strings.Split(f, "\n")
	herds := &SeaCucumberGrid{
		Grid: make([][]rune, 0, len(rows)),
		Rows: len(rows),
		Cols: len(rows[0]),
	}
	for _, row := range rows {
		herds.Grid = append(herds.Grid, []rune(row))
	}
	return herds
}
