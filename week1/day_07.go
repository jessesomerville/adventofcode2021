package week1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func WhaleVsCrabs(f string) int {
	posStr := strings.Split(f, ",")
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

// I have no idea why its one more or less than the average sometimes.  This is the distribution:
//    Less than avg:  00.505%
//    Average:        99.055%
//    More than avg:  00.440%
func WhaleVsCrabsGas(f string) int {
	posStr := strings.Split(f, ",")
	positions := make([]int, len(posStr))
	sum := 0
	for i, s := range posStr {
		positions[i], _ = strconv.Atoi(s)
		sum += positions[i]
	}

	avg := int(math.Round(float64(sum) / float64(len(positions))))

	var minGas1, minGas2, minGas3 int
	for _, p := range positions {
		dist1 := math.Abs(float64(p - avg))
		minGas1 += int((dist1 * (dist1 + 1)) / 2)
		dist2 := math.Abs(float64(p - avg + 1))
		minGas2 += int((dist2 * (dist2 + 1)) / 2)
		dist3 := math.Abs(float64(p - avg - 1))
		minGas3 += int((dist3 * (dist3 + 1)) / 2)
	}

	min := minGas1
	if minGas2 < minGas1 {
		min = minGas2
	}
	if minGas3 < min {
		return minGas3
	}
	return min
}
