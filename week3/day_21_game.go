package week3

// Player represents a single Dirac player.
type Player struct {
	Position int
	Score    int
}

// NewPlayers returns a new set of Dirac players.
func NewPlayers(pos1, pos2 int) []*Player {
	return []*Player{{Position: pos1}, {Position: pos2}}
}

// D100 represents the deterministic 100 sided die.
type D100 struct {
	TimesRolled int
	NextRoll    int
}

// Roll3 rolls the D100 three times and returns the sum.
func (d *D100) Roll3() int {
	return d.Roll() + d.Roll() + d.Roll()
}

// Roll rolls the D100 once.
func (d *D100) Roll() int {
	rollVal := (d.NextRoll % 100) + 1
	d.NextRoll++
	d.TimesRolled++
	return rollVal
}

// WinPaths counts the number of ways to win given any number of turns.
type WinPaths struct {
	Wins map[int]int
}

// Calculate recursively finds win conditions.
func (wp *WinPaths) Calculate(turn, score, pos int) {
	if score >= 21 {
		wp.Wins[turn-1] += (turn - 1) * 3
		return
	}

	for i := 3; i <= 9; i++ {
		newPos := (pos+i-1)%10 + 1
		wp.Calculate(turn+1, score+newPos, newPos)
	}
}
