package common

import "sync"

type Stack struct {
	Mut   sync.Mutex
	Stack []rune
}

func NewStack() *Stack {
	return &Stack{
		Mut:   sync.Mutex{},
		Stack: make([]rune, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.Stack)
}

func (s *Stack) Push(v rune) {
	s.Mut.Lock()
	defer s.Mut.Unlock()
	s.Stack = append(s.Stack, v)
}

func (s *Stack) Pop() (rune, bool) {
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
