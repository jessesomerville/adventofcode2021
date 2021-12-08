package main

import (
	"fmt"
	"math"
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

func sevenSegmentDecode() {
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
	x := inputs[1].signals

	sort.Slice(x, func(i, j int) bool {
		return len(x[i]) < len(x[j])
	})

	sum := 0
	// reverse i
	for i, o := range inputs[1].output {
		if len(o) == 2 {
			sum += int(math.Pow10(i))
		} else if len(o) == 3 {
			sum += int(7 * math.Pow10(i))
		} else if len(o) == 4 {
			sum += int(4 * math.Pow10(i))
		} else if len(o) == 7 {
			sum += int(8 * math.Pow10(i))
		}
	}
	fmt.Println(sum)

	// one := x[0]
	// seven := x[1]
	// four := x[2]
	// eight := x[9]
	// fiveLength := x[3:6]
	// sixLength := x[6:9]

}

func removeItem(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

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
// func sevenSegmentDecode() {
// 	// 	1: 2,
// 	// 	7: 3,
// 	// 	4: 4,
// 	// 	5: 5,
// 	// 	2: 5,
// 	// 	3: 5,
// 	// 	0: 6,
// 	// 	6: 6,
// 	// 	9: 6,
// 	// 	8: 7,

// 	inputs := parseInputSegments(testSegments)
// 	x := inputs[0].signals

// 	toFind := map[string]bool{}
// 	for _, i := range inputs[0].output {
// 		toFind[i] = true
// 	}

// 	sort.Slice(x, func(i, j int) bool {
// 		return len(x[i]) < len(x[j])
// 	})
// 	one := x[0]
// 	// seven := x[1]
// 	// four := x[2]
// 	fiveLength := x[3:6]
// 	sixLength := x[6:9]

// 	three := ""
// 	idx := 0
// 	for _, x := range fiveLength {
// 		if strings.ContainsRune(x, rune(one[0])) && strings.ContainsRune(x, rune(one[1])) {
// 			three = x
// 			break
// 		}
// 		idx++
// 	}
// 	fiveLength = append(fiveLength[:idx], fiveLength[idx+1:]...)
// 	six := ""
// 	idx = 0
// 	for _, x := range sixLength {
// 		if strings.ContainsRune(x, rune(one[0])) != strings.ContainsRune(x, rune(one[1])) {
// 			six = x
// 			break
// 		}
// 		idx++
// 	}
// 	sixLength = append(sixLength[:idx], sixLength[idx+1:]...)
// 	// // eight := x[8]

// }

// func removeItem(s []string, i int) []string {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
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
