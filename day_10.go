package main

import (
	_ "embed"
	"sort"
	"strings"
	"sync"
)

var (
	//go:embed inputs/day_10.txt
	chunksFile string

	matchLookup = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
)

func syntaxScoring() int {
	scoreLookup := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	bracketStack := newStack()

	totalPoints := 0
Loop:
	for _, chunk := range strings.Split(chunksFile, "\n") {
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

func syntaxScoringIncomplete() int {
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
	for _, chunk := range strings.Split(chunksFile, "\n") {
		totalPoints := 0
		bracketStack := newStack()
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

type stack struct {
	Mut   sync.Mutex
	Stack []rune
}

func newStack() *stack {
	return &stack{
		sync.Mutex{},
		make([]rune, 0),
	}
}

func (s *stack) Push(v rune) {
	s.Mut.Lock()
	defer s.Mut.Unlock()
	s.Stack = append(s.Stack, v)
}

func (s *stack) Pop() (rune, bool) {
	s.Mut.Lock()
	defer s.Mut.Unlock()

	l := len(s.Stack)
	if l == 0 {
		return '0', false
	}
	val := s.Stack[l-1]
	s.Stack = s.Stack[:l-1]
	return val, true
}
