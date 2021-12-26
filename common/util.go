package common

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// JoinFloat64 concatenates the elements of its first argument to create a string delimeted by sep.
func JoinFloat64(elems []float64, sep string) string {
	converted := make([]string, 0, len(elems))
	for _, e := range elems {
		elemStr := fmt.Sprintf("%.1f", e)
		converted = append(converted, elemStr)
	}
	return strings.Join(converted, sep)
}

// SortUniq returns a sorted unique set of the input ints.
func SortUniq(nums ...int) []int {
	numMap := make(map[int]bool)
	for _, n := range nums {
		numMap[n] = true
	}
	uniqNums := make([]int, 0, len(numMap))
	for n := range numMap {
		uniqNums = append(uniqNums, n)
	}
	sort.Ints(uniqNums)
	return uniqNums
}

func Max(nums ...int) int {
	var maxNum int
	for i, n := range nums {
		if i == 0 {
			maxNum = n
			continue
		}
		if n > maxNum {
			maxNum = n
		}
	}
	return maxNum
}

func Min(nums ...int) int {
	var minNum int
	for i, n := range nums {
		if i == 0 {
			minNum = n
			continue
		}
		if n < minNum {
			minNum = n
		}
	}
	return minNum
}

func DistBetween(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
