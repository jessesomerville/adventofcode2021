package week2

import (
	"sort"
	"strings"

	"github.com/jessesomerville/adventofcode2021/common"
)

var (
	matchLookup = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
)

func SyntaxScoring(f string) int {
	scoreLookup := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	bracketStack := common.NewStack()

	totalPoints := 0
Loop:
	for _, chunk := range strings.Split(f, "\n") {
		for _, c := range chunk {
			switch c {
			case '(', '[', '{', '<':
				bracketStack.Push(c)
			case ')', ']', '}', '>':
				val, ok := bracketStack.Pop()
				if !ok {
					totalPoints += scoreLookup[c]
					continue Loop
				}
				if val != matchLookup[c] {
					totalPoints += scoreLookup[c]
					continue Loop
				}
			}
		}
	}
	return totalPoints
}

func SyntaxScoringIncomplete(f string) int {
	scoreLookup := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	closeLookup := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	allScores := []int{}
Loop:
	for _, chunk := range strings.Split(f, "\n") {
		totalPoints := 0
		bracketStack := common.NewStack()
		for _, c := range chunk {
			switch c {
			case '(', '[', '{', '<':
				bracketStack.Push(c)
			case ')', ']', '}', '>':
				val, ok := bracketStack.Pop()
				if !ok {
					continue Loop
				}
				if val != matchLookup[c] {
					continue Loop
				}
			}
		}
		val, ok := bracketStack.Pop()
		for ok {
			totalPoints *= 5
			totalPoints += scoreLookup[closeLookup[val]]
			val, ok = bracketStack.Pop()
		}
		allScores = append(allScores, totalPoints)
	}

	sort.Ints(allScores)
	return allScores[len(allScores)/2]
}
