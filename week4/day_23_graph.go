package week4

// Types of spaces within the burrow.
const (
	AmberRoom = iota
	BronzeRoom
	CopperRoom
	DesertRoom
	Hallway
	OutsideOfRoom
)

// Amphipod represents one of the 8 amphipods.
type Amphipod struct {
	Type     int
	Energy   int
	Position *Space
}

// Space represents a single burrow space.
type Space struct {
	ID        int
	Type      int
	Neighbors []*Neighbor
}

// Neighbor represents the bordering spaces.
type Neighbor struct {
	*Space
	Blocked bool
}

// UnblockSpace adds a connection from s to n.
func (s *Space) UnblockSpace(n *Space) {
	for _, neighbor := range s.Neighbors {
		if neighbor.Space == n {
			neighbor.Blocked = false
			return
		}
	}
}

// BlockSpace removes a connection from s to n.
func (s *Space) BlockSpace(n *Space) {
	for _, neighbor := range s.Neighbors {
		if neighbor.Space == n {
			neighbor.Blocked = true
			return
		}
	}
}

// Burrow represents the graph of the burrow spaces.
type Burrow struct {
	Spaces []*Space
}

// NewBurrow creates a new empty burrow.
func NewBurrow() *Burrow {
	spaces := make([]*Space, 0, 19)
	for spaceLoc := 0; spaceLoc < 19; spaceLoc++ {
		newSpace := &Space{
			ID:        spaceLoc,
			Neighbors: make([]*Neighbor, 0, 3),
		}
		switch spaceLoc {
		case 0, 1, 3, 5, 7, 9, 10:
			newSpace.Type = Hallway
		case 2, 4, 6, 8:
			newSpace.Type = OutsideOfRoom
		case 11, 12:
			newSpace.Type = AmberRoom
		case 13, 14:
			newSpace.Type = BronzeRoom
		case 15, 16:
			newSpace.Type = CopperRoom
		case 17, 18:
			newSpace.Type = DesertRoom
		}
		spaces = append(spaces, newSpace)
	}
	b := &Burrow{spaces}
	for _, edge := range BurrowEdges {
		s1 := b.Spaces[edge[0]]
		s2 := b.Spaces[edge[1]]
		s1.Neighbors = append(s1.Neighbors, &Neighbor{s2, false})
		s2.Neighbors = append(s2.Neighbors, &Neighbor{s1, false})
	}
	return b
}

// NewAmphipod returns a new amphipod.
func (b *Burrow) NewAmphipod(t, location int) *Amphipod {
	occupied := b.Spaces[location]
	for _, n := range occupied.Neighbors {
		n.BlockSpace(occupied)
	}

	return &Amphipod{
		Type:     t,
		Energy:   EnergyByType[t],
		Position: occupied,
	}
}

// MoveAmphipod places the amphipod in a new space.
func (b *Burrow) MoveAmphipod(a *Amphipod, s *Space) {
	currSpace := a.Position
	for _, n := range currSpace.Neighbors {
		n.UnblockSpace(currSpace)
	}
	for _, n := range s.Neighbors {
		n.BlockSpace(s)
	}
	a.Position = s
}
