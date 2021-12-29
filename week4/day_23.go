package week4

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
	// burrow := NewBurrow()
	// amber1, amber2 := placeAmphipods(burrow, Amber, burrowTest[Amber])
	// bronze1, bronze2 := placeAmphipods(burrow, Bronze, burrowTest[Bronze])
	// copper1, copper2 := placeAmphipods(burrow, Copper, burrowTest[Copper])
	// desert1, desert2 := placeAmphipods(burrow, Desert, burrowTest[Desert])
	return 0
}

func placeAmphipods(b *Burrow, t int, locs []int) (a1, a2 *Amphipod) {
	a1 = b.NewAmphipod(t, locs[0])
	a2 = b.NewAmphipod(t, locs[1])
	return a1, a2
}

// 62568 too high
