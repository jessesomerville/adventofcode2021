package week3

import (
	"fmt"
	"math"
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
	wp1 := &WinPaths{
		Wins: make(map[int]int),
	}
	wp1.Calculate(1, 0, 4)

	wp2 := &WinPaths{
		Wins: make(map[int]int),
	}
	wp2.Calculate(1, 0, 8)

	p2NoWin := make(map[int]int, len(wp2.Wins))

	for i := 3; i <= 10; i++ {
		didntWin := 0
		for j := i; j <= 10; j++ {
			didntWin += wp2.Wins[j]
		}
		p2NoWin[i] = didntWin
	}

	fmt.Println(math.Pow(3, 6))

	// p1Wins := 0
	// for turn, winCount := range wp1.Wins {
	// 	p1Wins += winCount * p2NoWin[turn]
	// }
	// fmt.Println(p1Wins)

	return 0
}
