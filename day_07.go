package main

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed inputs/day_07.txt
	crabsFile string
)

func whaleVsCrabs() int {
	posStr := strings.Split(crabsFile, ",")
	positions := make([]int, len(posStr))
	for i, s := range posStr {
		positions[i], _ = strconv.Atoi(s)
	}

	sort.Ints(positions)

	midPoint := len(positions) / 2
	mid1 := positions[midPoint]
	if len(positions)%2 == 0 {
		moves := 0
		for _, p := range positions {
			moves += int(math.Abs(float64(p - mid1)))
		}
		return moves
	}
	mid2 := positions[midPoint-1]
	var moves1, moves2 int
	for _, p := range positions {
		moves1 += int(math.Abs(float64(p - mid1)))
		moves2 += int(math.Abs(float64(p - mid2)))
	}
	if moves1 < moves2 {
		return moves1
	}
	return moves2
}

func whaleVsCrabsGas() int {
	posStr := strings.Split(crabsFile, ",")
	positions := make([]int, len(posStr))
	sum := 0
	for i, s := range posStr {
		positions[i], _ = strconv.Atoi(s)
		sum += positions[i]
	}

	avg := int(math.Round(float64(sum) / float64(len(positions))))
	var minGas int

	for _, p := range positions {
		dist := math.Abs(float64(p - avg))
		minGas += int((dist * (dist + 1)) / 2)
	}
	prevMin := 0
	i := 1
	for prevMin != minGas {
		target1 := avg + i
		target2 := avg - i
		var gas1, gas2 int
		for _, p := range positions {
			dist1 := math.Abs(float64(p - target1))
			dist2 := math.Abs(float64(p - target2))
			gas1 += int((dist1 * (dist1 + 1)) / 2)
			gas2 += int((dist2 * (dist2 + 1)) / 2)
		}

		prevMin = minGas
		if gas1 < minGas {
			minGas = gas1
		}
		if gas2 < minGas {
			minGas = gas2
		}
		i++
	}
	return minGas

}
