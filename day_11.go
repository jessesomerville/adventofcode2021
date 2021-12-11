package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

var (
	//go:embed inputs/day_11.txt
	octosFile  string
	octosWidth = 10
)

func getOctos(f string) []int {
	octoGrid := make([]int, octosWidth*octosWidth)
	for row, line := range strings.Split(f, "\n") {
		for col, c := range line {
			coord := octosWidth*row + col
			octoGrid[coord], _ = strconv.Atoi(string(c))
		}
	}
	return octoGrid
}

func dumboOctopus() int {
	octos := getOctos(octosFile)
	allPops := 0
	for x := 0; x < 100; x++ {
		var readyToPop []int
		for i := range octos {
			if octos[i] == 9 {
				octos[i] = 0
				readyToPop = append(readyToPop, i)
			} else {
				octos[i]++
			}
		}
		for _, i := range readyToPop {
			octos, allPops = popOctos(i, octos, allPops)
		}
	}
	return allPops
}

func dumboOctopusSync() int {
	octos := getOctos(octosFile)
	allPops := 0
	step := 1
	for {
		var readyToPop []int
		for i := range octos {
			if octos[i] == 9 {
				octos[i] = 0
				readyToPop = append(readyToPop, i)
			} else {
				octos[i]++
			}
		}
		prevPops := allPops
		for _, i := range readyToPop {
			octos, allPops = popOctos(i, octos, allPops)
		}
		if allPops-prevPops == 100 {
			return step
		}
		step++
	}
}

func popOctos(idx int, octos []int, popped int) ([]int, int) {
	popped++
	adj := adjacent(idx, octos)
	for _, a := range adj {
		if octos[a] == 0 {
			continue
		}
		if octos[a] == 9 {
			octos[a] = 0
			octos, popped = popOctos(a, octos, popped)
		} else {
			octos[a]++
		}
	}
	return octos, popped
}

func printOctos(octos []int) {
	for i, a := range octos {
		fmt.Print(a)
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
}

func adjacent(i int, octos []int) []int {
	adj := make([]int, 0, 8)
	row := i / octosWidth
	col := i % octosWidth
	up := octosWidth*(row-1) + col
	down := octosWidth*(row+1) + col
	right := i + 1
	left := i - 1
	isR := right < len(points) && right/octosWidth == row
	isL := left >= 0 && left/octosWidth == row
	if up >= 0 {
		adj = append(adj, up)
		if isL {
			adj = append(adj, up-1)
		}
		if isR {
			adj = append(adj, up+1)
		}
	}
	if down < len(octos) {
		adj = append(adj, down)
		if isL {
			adj = append(adj, down-1)
		}
		if isR {
			adj = append(adj, down+1)
		}
	}
	if isL {
		adj = append(adj, left)
	}
	if isR {
		adj = append(adj, right)
	}
	return adj
}
