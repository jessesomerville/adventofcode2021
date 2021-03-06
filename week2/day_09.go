package week2

import (
	"sort"
	"strconv"
	"sync"
)

var (
	mapWidth = 100
	points   []int

	once sync.Once
)

func getPoints(f string) []int {
	once.Do(func() {
		points = make([]int, len(f))
		for i, s := range f {
			points[i], _ = strconv.Atoi(string(s))
		}
	})
	return points
}

// Change to binary search
func SmokeBasin(f string) int {
	points = getPoints(f)
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

func SmokeBasinLargest(f string) int {
	points = getPoints(f)

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
		allBasins = append(allBasins, floodSearch(i, map[int]bool{}))
	}
	sort.Slice(allBasins, func(i, j int) bool {
		return allBasins[i] > allBasins[j]
	})
	return allBasins[0] * allBasins[1] * allBasins[2]
}

func floodSearch(idx int, searched map[int]bool) int {
	currPoint := points[idx]
	basin := 1
	searched[idx] = true
	row := idx / mapWidth
	col := idx % mapWidth
	up := mapWidth*(row-1) + col
	down := mapWidth*(row+1) + col
	left := idx + 1
	right := idx - 1
	if up >= 0 && !searched[up] {
		if points[up] > currPoint && points[up] != 9 {
			basin += floodSearch(up, searched)
		}
	}
	if down < len(points) && !searched[down] {
		if points[down] > currPoint && points[down] != 9 {
			basin += floodSearch(down, searched)
		}
	}
	if left < len(points) && left/mapWidth == row && !searched[left] {
		if points[left] > currPoint && points[left] != 9 {
			basin += floodSearch(left, searched)
		}
	}
	if right >= 0 && right/mapWidth == row && !searched[right] {
		if points[right] > currPoint && points[right] != 9 {
			basin += floodSearch(right, searched)
		}
	}
	return basin
}
