package week1

import (
	"strconv"
	"strings"
)

func LanternFish(f string) int {
	fish := make([]int, 9)
	for _, f := range strings.Split(f, ",") {
		idx, _ := strconv.Atoi(f)
		fish[idx]++
	}

	for day := 0; day < 256; day++ {
		newFish := make([]int, 9)
		tmpFish := fish[0]
		for i := 7; i >= 0; i-- {
			newFish[i] = fish[i+1]
		}
		newFish[6] += tmpFish
		newFish[8] += tmpFish
		fish = newFish
	}
	totalFish := 0
	for _, i := range fish {
		totalFish += i
	}
	return totalFish
}
