package main

import (
	_ "embed"
)

var (
	//go:embed inputs/day_23.txt
	amphipodFile string

	amphipodFileTest = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
)

func amphipod() int {
	return 0
}
