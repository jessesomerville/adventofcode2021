package main

import (
	"math"
	"sort"
	"strings"

	_ "embed"
)

//go:embed inputs/day_08.txt
var segmentsFile string

func sevenSegment() int {
	var segmentCounts = map[int]bool{
		2: true,
		4: true,
		3: true,
		7: true,
	}
	segments := parseInputSegments(segmentsFile)
	counts := 0
	for _, s := range segments {
		for _, x := range s.output {
			if segmentCounts[len(x)] {
				counts++
			}
		}
	}
	return counts
}

func sevenSegmentDecode() int {
	sumOfAll := 0
	for _, row := range parseInputSegments(segmentsFile) {
		x := row.signals

		sort.Slice(x, func(i, j int) bool {
			return len(x[i]) < len(x[j])
		})

		toFind := make([]string, 4)

		sum := 0
		for i, o := range row.output {
			idx := (len(row.output) - 1) - i
			if len(o) == 2 {
				sum += int(math.Pow10(idx))
			} else if len(o) == 3 {
				sum += int(7 * math.Pow10(idx))
			} else if len(o) == 4 {
				sum += int(4 * math.Pow10(idx))
			} else if len(o) == 7 {
				sum += int(8 * math.Pow10(idx))
			} else {
				toFind[idx] = o
			}
		}

	Loop:
		for idx, currNum := range toFind {
			if len(currNum) == 0 {
				continue
			}
			if len(currNum) == 5 {
				three := findThree(x[0], x[3:6])

				if equalStrings(three, currNum) {
					sum += int(3 * math.Pow10(idx))
					continue
				}
				six := findSix(x[0], x[6:9])
				for _, seg := range currNum {
					if !strings.ContainsRune(six, seg) {
						sum += int(2 * math.Pow10(idx))
						continue Loop
					}
				}
				sum += int(5 * math.Pow10(idx))
			} else if len(currNum) == 6 {
				if strings.ContainsRune(currNum, rune(x[0][0])) != strings.ContainsRune(currNum, rune(x[0][1])) {
					sum += int(6 * math.Pow10(idx))
					continue
				}
				three := findThree(x[0], x[3:6])
				for _, seg := range three {
					if !strings.ContainsRune(currNum, seg) {
						// 0
						continue Loop
					}
				}
				sum += int(9 * math.Pow10(idx))
			}
		}
		sumOfAll += sum
	}
	return sumOfAll
}

func findThree(one string, fiveLen []string) string {
	for _, n := range fiveLen {
		if strings.ContainsRune(n, rune(one[0])) && strings.ContainsRune(n, rune(one[1])) {
			return n
		}
	}
	return ""
}

func findSix(one string, sixLen []string) string {
	for _, n := range sixLen {
		if strings.ContainsRune(n, rune(one[0])) != strings.ContainsRune(n, rune(one[1])) {
			return n
		}
	}
	return ""
}

func equalStrings(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	am := make(map[rune]bool)
	for _, i := range a {
		am[i] = true
	}
	for _, i := range b {
		if !am[i] {
			return false
		}
	}
	return true
}

type segment struct {
	signals []string
	output  []string
}

func parseInputSegments(in string) []*segment {
	ts := strings.Split(in, "\n")
	ret := make([]*segment, len(ts))
	for i, s := range ts {
		x := strings.Split(s, " | ")
		ret[i] = &segment{
			signals: strings.Split(x[0], " "),
			output:  strings.Split(x[1], " "),
		}
	}
	return ret
}
