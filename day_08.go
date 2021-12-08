package main

import (
	"fmt"
	"sort"
	"strings"

	_ "embed"
)

//go:embed inputs/day_08.txt
var segmentsFile string

var testSegments = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func sevenSegment() {
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
	fmt.Println(counts)
}

// d3 d2 d1 d0
// 0  0  0  0
// 0  0  0  1
// 0  0  1  0

// a, b    : (d1 && d0)
// a, d, e : (!d2 && !d0)
// b, f    : (!d1 && !d0)
// d, e, g : (d1 && !d0)
// d, g    : (!d2 && d1)

// if a then d3 or d2 & d0 or d1 & d0 or !d2 & !d0

// func kTable() {
// 	var d3, d2, d1, d0 bool
// 	a := d3 || (d2 && d0) || (d1 && d0) || (!d2 && !d0)
// 	b := d3 || (d1 && d0) || (!d1 && !d0) || !d2
// 	c := d3 || d2 || !d1 || d0
// 	d := (d2 && !d1 && d0) || (!d2 && d1) || (d1 && !d0) || (!d2 && !d0)
// 	e := (d1 && !d0) || (!d2 && !d0)
// 	f := d3 || (!d1 && !d0) || (d2 && !d1) || (d2 && !d0)
// 	g := d3 || (d2 && !d1) || (d1 && !d0) || (!d2 && d1)

// }

//  aaaa
// f    b
// f    b
//  gggg
// e    c
// e    c
//  dddd

//  1111
// 2    3
// 2    3
//  4444
// 5    6
// 5    6
//  7777
func sevenSegmentDecode() {
	// segCount := map[int]int{
	// 	0: 6,
	// 	1: 2,
	// 	2: 5,
	// 	3: 5,
	// 	4: 4,
	// 	5: 5,
	// 	6: 6,
	// 	7: 3,
	// 	8: 7,
	// 	9: 6,
	// }
	// 	1: 2,
	// 	7: 3,
	// 	4: 4,
	// 	5: 5,
	// 	2: 5,
	// 	3: 5,
	// 	0: 6,
	// 	6: 6,
	// 	9: 6,
	// 	8: 7,

	inputs := parseInputSegments(testSegments)
	x := inputs[0].signals
	// segMaps := make([]map[rune]bool, 8)
	// for i, s := range x {
	// 	m := make(map[rune]bool, len(s))
	// 	for _, r := range s {
	// 		m[r] = true
	// 	}
	// 	segMaps[i] = m
	// }
	sort.Slice(x, func(i, j int) bool {
		return len(x[i]) < len(x[j])
	})
	// one := x[0]
	// seven := x[1]
	// four := x[2]
	fiveLength := x[3:6]
	sixLength := x[6:9]

	fmt.Println(fiveLength)
	fmt.Println(sixLength)
	// // eight := x[8]

	// a := getDiff(one, seven)[0]

	// 6 is six length that doesn't contain all of 1
	// 9 shares g with 6, 0 doesn't

}

// b should contain more segments
// func getDiff(a, b string) []rune {
// 	am := make(map[rune]bool)
// 	for _, i := range a {
// 		am[i] = true
// 	}
// 	diff := []rune{}
// 	for _, i := range b {
// 		if !am[i] {
// 			diff = append(diff, i)
// 		}
// 	}
// 	return diff
// }

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
