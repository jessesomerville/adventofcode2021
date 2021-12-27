package week4

// Amphipod represents one of the amphipods in the burrow.
type Amphipod struct {
	Type     int
	Location *Space
	State    int
}

// Space represents a single space in the burrow.
type Space struct {
	Type       int
	Accessible map[int]bool
}

// Burrow represents the burrow as a whole.
type Burrow struct {
	Amphipods  map[int][]*Amphipod
	Spaces     []*Space
	DistMatrix [][]int
}

// NewBurrow returns a burrow with the specified initial state.
func NewBurrow(initState map[int][]int) *Burrow {
	b := &Burrow{
		Amphipods:  make(map[int][]*Amphipod, 4),
		Spaces:     make([]*Space, 0, 19),
		DistMatrix: DistanceMatrix,
	}
	for i := 0; i < 19; i++ {
		newRoom := &Space{Accessible: initialAccess(i)}
		switch i {
		case 0, 1, 3, 5, 7, 9, 10:
			newRoom.Type = Hallway
		case 2, 4, 6, 8:
			newRoom.Type = OutsideOfRoom
		case 11, 12:
			newRoom.Type = AmberRoom
		case 13, 14:
			newRoom.Type = BronzeRoom
		case 15, 16:
			newRoom.Type = CopperRoom
		case 17, 18:
			newRoom.Type = DesertRoom
		}
		b.Spaces = append(b.Spaces, newRoom)
	}
	for aType, pos := range initState {
		thisType := make([]*Amphipod, 0, 2)
		for _, position := range pos {
			amph := &Amphipod{
				Type:     aType,
				Location: b.Spaces[position],
				State:    Unmoved,
			}
			thisType = append(thisType, amph)
		}
		b.Amphipods[aType] = thisType
	}
	return b
}

func initialAccess(idx int) map[int]bool {
	acc := make(map[int]bool, 19)
	if idx == 12 || idx == 14 || idx == 16 || idx == 18 {
		return acc
	}
	for i := 0; i < 11; i++ {
		acc[i] = true
	}
	acc[idx] = false
	return acc
}
