package week1

import (
	"strconv"
	"strings"
)

func SonarSweep(f string) int {
	answer := 0
	d := strings.Split(f, "\n")
	for i := 0; i < len(d)-1; i++ {
		dx, _ := strconv.Atoi(d[i])
		dy, _ := strconv.Atoi(d[i+1])
		if dx < dy {
			answer++
		}
	}
	return answer
}

func SonarSweepSlidingWindow(f string) int {
	answer := 0
	d := strings.Split(f, "\n")
	for i := 0; i < len(d)-3; i++ {
		dx, _ := strconv.Atoi(d[i])
		dy, _ := strconv.Atoi(d[i+3])
		if dx < dy {
			answer++
		}
	}
	return answer
}
