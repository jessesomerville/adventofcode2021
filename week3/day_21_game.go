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

var (
	// RollDistribution represents the number of ways (values) to roll a number (keys) by rolling a D3 three times.
	RollDistribution = map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	// WinsFromPosition maps the current position to the points needed to win and the number of ways to get there.
	WinsFromPosition = make(map[int]map[int]int, 10)
)

func positionsToWinConditions() {
	for pos := 1; pos <= 10; pos++ {
		ptsFromPos := make(map[int]int, 7)
		for roll := 3; roll <= 9; roll++ {
			newPos := (pos+roll-1)%10 + 1
			ptsFromPos[newPos] = RollDistribution[roll]
		}
		WinsFromPosition[pos] = ptsFromPos
	}
}

// WinCondition is the number of points needed to win the game.
const WinCondition = 21

// PlayerOutcomes represents the number of ways a player could win at a given round.
type PlayerOutcomes struct {
	WinsByTurn map[int]int
}

// Simulate finds all of the possible ways a player can win.
func (po *PlayerOutcomes) Simulate(turn, position, score, convergences int) {
	for rollSum, waysToRoll := range RollDistribution {
		newPos := (position+rollSum-1)%10 + 1
		newScore := score + newPos
		if newScore >= WinCondition {
			po.WinsByTurn[turn] += convergences * waysToRoll
		} else {
			po.Simulate(turn+1, newPos, newScore, convergences*waysToRoll)
		}
	}
}
