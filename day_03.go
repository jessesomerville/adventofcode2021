package main

import (
	_ "embed"
	"math/big"
)

//go:embed inputs/day_03.txt
var diagFile string

func binaryDiagnostic() int64 {
	gamma := new(big.Int)
	epsilon := new(big.Int).SetInt64((1 << 12) - 1)
	one := make([]int, 12)

	for line := range parseFile(diagFile) {
		n, _ := new(big.Int).SetString(line, 2)
		for i := n.BitLen() - 1; i >= 0; i-- {
			if gamma.Bit(i) == 1 {
				continue
			}
			if n.Bit(i) == 1 {
				one[i]++
			}
			if one[i] >= 500 {
				gamma.SetBit(gamma, i, 1)
				epsilon.SetBit(epsilon, i, 0)
			}
		}
	}
	return gamma.Mul(gamma, epsilon).Int64()
}

func binaryDiagnosticLifeSupport() int64 {
	oxy := new(big.Int)
	co2 := new(big.Int)
	tree, bitLen := makeTrie()
	t1, t2 := tree, tree

	for i := bitLen - 1; i >= 0; i-- {
		if t1.zero.weight > t1.one.weight {
			t1 = t1.zero
		} else {
			oxy.SetBit(oxy, i, 1)
			t1 = t1.one
		}

		if t2.zero.weight <= t2.one.weight && t2.zero.weight != 0 {
			t2 = t2.zero
		} else {
			co2.SetBit(co2, i, 1)
			t2 = t2.one
		}
	}
	return oxy.Mul(oxy, co2).Int64()
}

type node struct {
	zero, one *node
	weight    int
}

func newNode() *node {
	return &node{
		zero: &node{weight: 0},
		one:  &node{weight: 0},
	}
}

func makeTrie() (*node, int) {
	root := newNode()
	bitLen := 0
	for n := range parseFile(diagFile) {
		bitLen = len(n)
		currParent := root
		for _, c := range n {
			switch c {
			case '0':
				if currParent.zero.weight == 0 {
					currParent.zero = newNode()
				}
				currParent.zero.weight++
				currParent = currParent.zero
			case '1':
				if currParent.one.weight == 0 {
					currParent.one = newNode()
				}
				currParent.one.weight++
				currParent = currParent.one
			}
		}
	}
	return root, bitLen
}
