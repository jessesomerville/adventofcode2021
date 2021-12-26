package week3

type player struct {
	ID       int
	Position int
	Score    int
}

var (
	diracPlayers = []*player{
		{
			ID:       1,
			Position: 1,
		},
		{
			ID:       2,
			Position: 2,
		},
	}
	diracPlayersTest = []*player{
		{
			ID:       1,
			Position: 4,
		},
		{
			ID:       2,
			Position: 8,
		},
	}
)

// D100 represents the deterministic 100 sided die.
type D100 struct {
	TimesRolled int
	NextRoll    int
}

// Roll rolls the D100 three times and returns the sum.
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

func DiracDice() int {
	players := diracPlayers
	die := &D100{
		TimesRolled: 0,
		NextRoll:    0,
	}

	turn := 0
	for players[0].Score < 1000 && players[1].Score < 1000 {
		currPlayer := players[turn%2]
		roll := die.Roll3()
		newPos := currPlayer.Position + roll - 1
		currPlayer.Position = (newPos % 10) + 1
		currPlayer.Score += currPlayer.Position
		turn++
	}

	var loserScore int
	if players[0].Score == 1000 {
		loserScore = players[1].Score
	} else {
		loserScore = players[0].Score
	}
	return loserScore * die.TimesRolled
}
