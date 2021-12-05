package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

var (
	//go:embed inputs/day_04.txt
	bingoFile string

	// 2d coords -> 1d coords = (width * row) + col
	// This is just for up diagonal because down diagonals have the same row and col.
	downDiagCoords = map[int]bool{
		4:  true,
		8:  true,
		16: true,
		20: true,
	}
)

type cell struct {
	row    int
	col    int
	marked bool
}

type board struct {
	values           map[string]*cell
	rows, cols       [5]int // The count of unmarked cells for each row and column.
	upDiag, downDiag int    // The count of unmarked cells for each diagonal.
	out              bool
}

func (b *board) won(draw string) bool {
	cell := b.values[draw]
	cell.marked = true
	b.rows[cell.row]--
	b.cols[cell.col]--
	if cell.row == cell.col {
		b.downDiag--
	}
	if downDiagCoords[5*cell.row+cell.col] {
		b.upDiag--
	}
	if b.rows[cell.row] == 0 || b.cols[cell.col] == 0 || b.downDiag == 0 || b.upDiag == 0 {
		b.out = true
		return true
	}
	return false
}

func (b *board) score(draw string) int {
	sum := 0
	for val, cell := range b.values {
		if !cell.marked {
			valInt, _ := strconv.Atoi(val)
			sum += valInt
		}
	}
	drawInt, _ := strconv.Atoi(draw)
	return drawInt * sum
}

// This uses a ridiculous amount of memory, but it's fast :shrug:
func parseBoards(f string) ([]string, map[string][]*board, int) {
	sections := strings.SplitN(f, "\n", 2)
	draws := strings.Split(sections[0], ",")

	boardsStr := strings.Split(sections[1], "\n\n")
	boards := make(map[string][]*board, len(boardsStr)) // map from values to the boards that have it.

	for _, b := range boardsStr {
		fields := strings.Fields(b)
		width := int(math.Sqrt(float64(len(fields))))
		thisBoard := &board{
			values:   make(map[string]*cell, len(fields)),
			rows:     [5]int{5, 5, 5, 5, 5},
			cols:     [5]int{5, 5, 5, 5, 5},
			upDiag:   5,
			downDiag: 5,
		}
		for i, f := range fields {
			thisBoard.values[f] = &cell{row: i / width, col: i % width}
			if v, ok := boards[f]; ok {
				boards[f] = append(v, thisBoard)
			} else {
				boards[f] = []*board{thisBoard}
			}
		}
	}
	return draws, boards, len(boardsStr)
}

func giantSquid() int {
	draws, boards, _ := parseBoards(bingoFile)
	for _, draw := range draws {
		if matchingBoards, ok := boards[draw]; ok {
			for _, b := range matchingBoards {
				if b.won(draw) {
					return b.score(draw)
				}
			}
		}
	}
	return 0
}

func giantSquidLastWinner() int {
	draws, boards, players := parseBoards(bingoFile)
	for _, draw := range draws {
		if matchingBoards, ok := boards[draw]; ok {
			for _, b := range matchingBoards {
				if b.out {
					continue
				}
				if b.won(draw) {
					if players == 1 {
						return b.score(draw)
					}
					players--
				}
			}
		}
	}
	return 0
}
