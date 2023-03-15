package week1

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
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

func (b *board) won(c *cell) bool {
	c.marked = true
	b.rows[c.row]--
	b.cols[c.col]--
	coords := 5*c.row + c.col // 2d coords -> 1d coords = (width * row) + col
	switch coords {
	case 2:
		b.downDiag--
		b.upDiag--
	case 0, 6, 12, 24:
		b.downDiag--
	case 4, 8, 16, 20:
		b.upDiag--
	}

	if b.rows[c.row] == 0 || b.cols[c.col] == 0 || b.downDiag == 0 || b.upDiag == 0 {
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

var fieldsPool = sync.Pool{
	New: func() interface{} {
		return make([]string, 25)
	},
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func parseBoards(f string) ([]string, []*board, int) {
	lines := strings.Split(f, "\n\n")
	draws := strings.Split(lines[0], ",")

	boards := make([]*board, len(lines)-1)
	for bi, s := range lines[1:] {
		fields := fieldsPool.Get().([]string)
		// The following ~20 lines is strings.Fields(), but I modified it to use a sync.Pool to save on memory.
		na := 0
		fieldStart := 0
		i := 0
		for i < len(s) && asciiSpace[s[i]] != 0 {
			i++
		}
		fieldStart = i
		for i < len(s) {
			if asciiSpace[s[i]] == 0 {
				i++
				continue
			}
			fields[na] = s[fieldStart:i]
			na++
			i++
			for i < len(s) && asciiSpace[s[i]] != 0 {
				i++
			}
			fieldStart = i
		}
		if fieldStart < len(s) {
			fields[na] = s[fieldStart:]
		}

		thisBoard := &board{
			values:   make(map[string]*cell, 25),
			rows:     [5]int{5, 5, 5, 5, 5},
			cols:     [5]int{5, 5, 5, 5, 5},
			upDiag:   5,
			downDiag: 5,
		}
		for fi, f := range fields {
			thisBoard.values[f] = &cell{row: fi / 5, col: fi % 5}
		}
		fieldsPool.Put(fields)
		boards[bi] = thisBoard
	}
	return draws, boards, len(lines[1:])
}

func GiantSquid(f string) int {
	draws, boards, _ := parseBoards(f)
	for _, draw := range draws {
		for _, b := range boards {
			if c, ok := b.values[draw]; ok {
				if b.won(c) {
					fmt.Println(b)
					return b.score(draw)
				}
			}
		}
	}
	return 0
}

func GiantSquidLastWinner(f string) int {
	draws, boards, _ := parseBoards(f)
	var maxWinner *board
	maxMoves := 0
	maxDraw := ""

	// Since we are eventually completing every board anyway, it's faster to just solve each board
	// and compare the number of moves it took.
	for _, b := range boards {
		for i, draw := range draws {
			if c, ok := b.values[draw]; ok {
				if b.won(c) {
					if i > maxMoves {
						maxMoves = i
						maxDraw = draw
						maxWinner = b
					}
					break
				}
			}
		}
	}
	return maxWinner.score(maxDraw)
}

// Print a bingo board for Day 04
func (b *board) String() string {
	rows := make([][]string, 5)

	for val, cell := range b.values {
		if len(rows[cell.row]) == 0 {
			rows[cell.row] = make([]string, 5)
		}
		rows[cell.row][cell.col] = val
	}

	buf := new(strings.Builder)
	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.AlignRight)
	for _, row := range rows {
		for _, cellVal := range row {
			thisCell := b.values[cellVal]
			if thisCell.marked {
				fmt.Fprintf(w, "\x1b[34m%s\x1b[0m\t", cellVal)
			} else {
				fmt.Fprintf(w, "\x1b[97m%s\x1b[0m\t", cellVal)
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buf.String()
}
