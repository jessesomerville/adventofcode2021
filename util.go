package main

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

func parseFile(f string) <-chan string {
	c := make(chan string)
	go func() {
		for _, line := range strings.Split(f, "\n") {
			c <- line
		}
		close(c)
	}()
	return c
}

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
				fmt.Fprintf(w, "%s\t", color.BlueString(cellVal))
			} else {
				fmt.Fprintf(w, "%s\t", color.HiWhiteString(cellVal))
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buf.String()
}
