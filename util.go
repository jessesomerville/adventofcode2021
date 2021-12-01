package main

import (
	"log"
	"strconv"
	"strings"
)

func allFileBytesToInts(f []byte) []int {
	lines := strings.Split(string(f), "\n")
	vals := make([]int, len(lines))

	for i, l := range lines {
		x, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		vals[i] = x
	}
	return vals
}
