package main

import (
	_ "embed"
	"fmt"
)

//go:embed inputs/day_01.txt
var depthsFile []byte

func sonarSweep() {
	d := allFileBytesToInts(depthsFile)
	answer := 0
	for i := 0; i < len(d)-1; i++ {
		if d[i] < d[i+1] {
			answer++
		}
	}
	fmt.Println(answer)
}

func sonarSweepSlidingWindow() {
	d := allFileBytesToInts(depthsFile)

	answer := 0
	for i := 0; i < len(d)-3; i++ {
		if d[i] < d[i+3] {
			answer++
		}
	}
	fmt.Println(answer)
}
