package main

import (
	_ "embed"
)

//go:embed inputs/day_03.txt
var diagFile string

func binaryDiagnostic() int {
	gamma := 0
	one := make([]int, 12)

	for line := range parseFile(diagFile) {
		for i, bit := range line {
			if one[i] >= 500 {
				continue
			}
			if bit == '1' {
				one[i]++
			}
			if one[i] >= 500 {
				gamma += 1 << (11 - i)
			}
		}
	}
	return gamma * (((1 << 12) - 1) - gamma)
}

func binaryDiagnosticLifeSupport() int {
	var oxygen, co2 int

	tree := makeTrie()
	t1, t2 := tree, tree

	for i := 0; i < 12; i++ {
		if t1.zero == nil {
			oxygen += 1 << (11 - i)
			t1 = t1.one
		} else if t1.one == nil {
			t1 = t1.zero
		} else if t1.one == nil || t1.zero.weight > t1.one.weight {
			t1 = t1.zero
		} else {
			oxygen += 1 << (11 - i)
			t1 = t1.one
		}

		if t2.zero == nil {
			co2 += 1 << (11 - i)
			t2 = t2.one
		} else if t2.one == nil {
			t2 = t2.zero
		} else if t2.zero.weight <= t2.one.weight && t2.zero.weight != 0 {
			t2 = t2.zero
		} else {
			co2 += 1 << (11 - i)
			t2 = t2.one
		}
	}

	return oxygen * co2
}

type node struct {
	zero, one *node
	weight    int
}

func makeTrie() *node {
	root := new(node)
	for n := range parseFile(diagFile) {
		currParent := root
		for _, c := range n {
			switch c {
			case '0':
				if currParent.zero == nil {
					currParent.zero = new(node)
				}
				currParent.zero.weight++
				currParent = currParent.zero
			case '1':
				if currParent.one == nil {
					currParent.one = new(node)
				}
				currParent.one.weight++
				currParent = currParent.one
			}
		}
	}
	return root
}
