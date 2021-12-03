package main

import (
	_ "embed"
	"math/big"
	"strings"
)

//go:embed inputs/day_03.txt
var diagFile string

func parseDiag() <-chan *big.Int {
	c := make(chan *big.Int)
	go func() {
		for _, line := range strings.Split(diagFile, "\n") {
			n, _ := new(big.Int).SetString(line, 2)
			c <- n
		}
		close(c)
	}()
	return c
}

func binaryDiagnostic() int64 {
	gamma := new(big.Int)
	epsilon := new(big.Int).SetInt64((1 << 12) - 1)
	one := make([]int, 12)

	for n := range parseDiag() {
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
	oxy := new(big.Int) // .SetInt64(getOxy())
	co2 := new(big.Int) // .SetInt64(getCo2())
	tree := makeTree()
	t1, t2 := tree, tree

	for i := 11; i >= 0; i-- {
		switch {
		case t1.one.weight < t1.zero.weight:
			oxy.SetBit(oxy, i, 0)
			t1 = t1.zero
		case t1.one.weight >= t1.zero.weight:
			oxy.SetBit(oxy, i, 1)
			t1 = t1.one
		}

		switch {
		case t2.zero.weight == 0:
			co2.SetBit(co2, i, 1)
			t2 = t2.one
		case t2.one.weight == 0:
			co2.SetBit(co2, i, 0)
			t2 = t2.zero
		case t2.zero.weight <= t2.one.weight:
			co2.SetBit(co2, i, 0)
			t2 = t2.zero
		case t2.zero.weight > t2.one.weight:
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

func makeTree() *node {
	root := newNode()

	for _, n := range strings.Split(diagFile, "\n") {
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
	return root
}
