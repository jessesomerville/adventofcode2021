package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day_01.txt
var depthsFile string

func sonarSweep() int {
	answer := 0
	d := strings.Split(depthsFile, "\n")
	for i := 0; i < len(d)-1; i++ {
		dx, _ := strconv.Atoi(d[i])
		dy, _ := strconv.Atoi(d[i+1])
		if dx < dy {
			answer++
		}
	}
	return answer
}

func sonarSweepSlidingWindow() int {
	answer := 0
	d := strings.Split(depthsFile, "\n")
	for i := 0; i < len(d)-3; i++ {
		dx, _ := strconv.Atoi(d[i])
		dy, _ := strconv.Atoi(d[i+3])
		if dx < dy {
			answer++
		}
	}
	return answer
}
