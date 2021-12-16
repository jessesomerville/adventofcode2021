package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/vinzmay/go-rope"
)

var (
	//go:embed inputs/day_14.txt
	polymerFile string

	polymerTest = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
)

func polymerization() int {
	template, rules := parsePolymer(polymerTest)

	for i := 0; i < 10; i++ {
		newTmpl := ""
		for idx := 0; idx < len(template)-1; idx++ {
			thisPiece := template[idx : idx+2]
			if insert, ok := rules[thisPiece]; ok {
				newTmpl += fmt.Sprintf("%s%s", insert, string(thisPiece[1]))
			}
		}
		template = string(template[0]) + newTmpl
	}
	runeCount := make(map[rune]int)
	for _, r := range template {
		if val, ok := runeCount[r]; ok {
			runeCount[r] = val + 1
		} else {
			runeCount[r] = 1
		}
	}

	max := -1
	min := 100000000000000000
	for _, v := range runeCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)

	return 0
}

func polymerization40() int {
	t, rules := parsePolymer(polymerTest)

	template := rope.New(t)

	fmt.Println(template)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for idx := 0; idx < template.Len()-1; idx++ {
			thisPiece := template.Substr(idx, 2)
			if insert, ok := rules[thisPiece.String()]; ok {
				template.Insert(idx+1, insert)
			}
		}
	}

	runeCount := make(map[rune]int)
	for _, r := range template.String() {
		if val, ok := runeCount[r]; ok {
			runeCount[r] = val + 1
		} else {
			runeCount[r] = 1
		}
	}

	max := -1
	min := 100000000000000000
	for _, v := range runeCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)

	return 0
}

func parsePolymer(f string) (string, map[string]string) {
	parts := strings.Split(f, "\n\n")
	template := parts[0]
	ruleStr := strings.Split(parts[1], "\n")

	rules := make(map[string]string, len(ruleStr))
	for _, r := range ruleStr {
		rs := strings.Split(r, " -> ")
		rules[rs[0]] = rs[1]
	}
	return template, rules
}
