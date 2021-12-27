package week3

import (
	"github.com/jessesomerville/adventofcode2021/common"
)

var (
	diracPlayers     = NewPlayers(1, 2)
	diracPlayersTest = NewPlayers(4, 8)
)

// Dirac is the solution for the practice round of Dirac.
func Dirac() int {
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

// DiracQuantum is the solution to Dirac with the quantum die.
func DiracQuantum() int {
	positionsToWinConditions()
	p1Games := &PlayerOutcomes{
		WinsByTurn: make(map[int]int),
	}
	p1Games.Simulate(1, 1, 0, 1)
	p2Games := &PlayerOutcomes{
		WinsByTurn: make(map[int]int),
	}
	p2Games.Simulate(1, 2, 0, 1)

	p1Divergences := 1
	p2Divergences := 1
	p1Wins := 0
	p2Wins := 0
	for i := 1; i <= 10; i++ {
		p1Wins += p1Games.WinsByTurn[i] * p1Divergences
		p1Divergences *= 27
		p1Divergences -= p2Games.WinsByTurn[i]

		p2Divergences *= 27
		p2Divergences -= p1Games.WinsByTurn[i]
		p2Wins += p2Games.WinsByTurn[i] * p2Divergences
	}

	return common.Max(p1Wins, p2Wins)
}
