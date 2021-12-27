package week4

import (
	"fmt"
	"strings"
)

var (
	amphipodFileTest = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

	burrowTest = map[int][]int{
		Amber:  {12, 18},
		Bronze: {11, 15},
		Copper: {13, 16},
		Desert: {14, 17},
	}
)

// SortBurrow rearranges the amphipods into their proper rooms.
func SortBurrow(f string) int {
	burrow := NewBurrow(burrowTest)
	buf := new(strings.Builder)
	for _, row := range burrow.DistMatrix {
		fmt.Fprint(buf, "{")
		for i, d := range row {
			if i != len(row)-1 {
				fmt.Fprintf(buf, "%d, ", d)
			} else {
				fmt.Fprintf(buf, "%d},\n", d)
			}
		}
	}
	fmt.Println(buf.String())
	return 0
}
