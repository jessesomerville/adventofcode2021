package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/day_09.txt
var heightMapFile string

var mapWidth = 100
var mapHeight = 100

// var mapWidth = 10
// var mapHeight = 5

// Change to binary search
func smokeBasin() int {
	lines := strings.Split(heightMapFile, "\n")
	points := make([]int, mapWidth*mapHeight)
	for ri, s := range lines {
		for ci, x := range s {
			coord := mapWidth*ri + ci
			points[coord], _ = strconv.Atoi(string(x))
		}
	}

	// printMap(points)
	riskSum := 0
	for i, currPoint := range points {
		row := i / mapWidth
		col := i % mapWidth
		up := mapWidth*(row-1) + col
		down := mapWidth*(row+1) + col
		left := i + 1
		right := i - 1
		if up >= 0 {
			if points[up] <= currPoint {
				continue
			}
		}
		if down < len(points) {
			if points[down] <= currPoint {
				continue
			}
		}
		if left < len(points) && left/mapWidth == row {
			if points[left] <= currPoint {
				continue
			}
		}
		if right >= 0 && right/mapWidth == row {
			if points[right] <= currPoint {
				continue
			}
		}
		riskSum += (currPoint + 1)
	}
	return riskSum
}

func smokeBasinLargest() int {
	lines := strings.Split(heightMapFile, "\n")
	points := make([]int, mapWidth*mapHeight)
	for ri, s := range lines {
		for ci, x := range s {
			coord := mapWidth*ri + ci
			points[coord], _ = strconv.Atoi(string(x))
		}
	}

	allBasins := []int{}
	for i, currPoint := range points {
		row := i / mapWidth
		col := i % mapWidth
		up := mapWidth*(row-1) + col
		down := mapWidth*(row+1) + col
		left := i + 1
		right := i - 1
		if up >= 0 {
			if points[up] <= currPoint {
				continue
			}
		}
		if down < len(points) {
			if points[down] <= currPoint {
				continue
			}
		}
		if left < len(points) && left/mapWidth == row {
			if points[left] <= currPoint {
				continue
			}
		}
		if right >= 0 && right/mapWidth == row {
			if points[right] <= currPoint {
				continue
			}
		}
		allBasins = append(allBasins, len(floodSearch(i, points, map[int]bool{})))
	}
	sort.Slice(allBasins, func(i, j int) bool {
		return allBasins[i] > allBasins[j]
	})
	return allBasins[0] * allBasins[1] * allBasins[2]
}

func floodSearch(idx int, points []int, searched map[int]bool) []int {
	currPoint := points[idx]
	basin := []int{currPoint}
	searched[idx] = true
	row := idx / mapWidth
	col := idx % mapWidth
	up := mapWidth*(row-1) + col
	down := mapWidth*(row+1) + col
	left := idx + 1
	right := idx - 1
	if up >= 0 && !searched[up] {
		if points[up] > currPoint && points[up] != 9 {
			basin = append(basin, floodSearch(up, points, searched)...)
		}
	}
	if down < len(points) && !searched[down] {
		if points[down] > currPoint && points[down] != 9 {
			basin = append(basin, floodSearch(down, points, searched)...)
		}
	}
	if left < len(points) && left/mapWidth == row && !searched[left] {
		if points[left] > currPoint && points[left] != 9 {
			basin = append(basin, floodSearch(left, points, searched)...)
		}
	}
	if right >= 0 && right/mapWidth == row && !searched[right] {
		if points[right] > currPoint && points[right] != 9 {
			basin = append(basin, floodSearch(right, points, searched)...)
		}
	}
	return basin
}
