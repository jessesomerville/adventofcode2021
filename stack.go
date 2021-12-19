package main

import "sync"

type stack struct {
	Mut   sync.Mutex
	Stack []rune
}

func newStack() *stack {
	return &stack{
		Mut:   sync.Mutex{},
		Stack: make([]rune, 0),
	}
}

func (s *stack) Len() int {
	return len(s.Stack)
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
