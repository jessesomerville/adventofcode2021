package week2

import (
	"fmt"
)

type Memoizer struct {
	Cache map[string][]rune
	Rules map[string]rune
}

func NewMemoizer(rules map[string]rune) *Memoizer {
	return &Memoizer{
		Cache: make(map[string][]rune),
		Rules: rules,
	}
}

func (m *Memoizer) Evaluate(s []rune) []rune {
	result := make([]rune, 0, len(s)*2-1)
	for i := 0; i < len(s); i += 16 {
		if i+16 >= len(s) {
			result = append(result, m.Insert(s[i:])...)
		} else {
			result = append(result, m.Insert(s[i:i+16])...)
			result = append(result, m.Lookup(s[i+15], s[i+16]))
		}
	}
	return result
}

func (m *Memoizer) Insert(s []rune) []rune {
	if res, ok := m.Cache[string(s)]; ok {
		return res
	}
	res := m.evalRules(s)
	m.Cache[string(s)] = res
	return res
}

func (m *Memoizer) evalRules(s []rune) []rune {
	res := make([]rune, 0, len(s)*2-1)
	for i := 0; i < len(s)-1; i++ {
		currR, nextR := s[i], s[i+1]
		res = append(res, currR)
		lookup := fmt.Sprintf("%s%s", string(currR), string(nextR))
		res = append(res, rune(m.Rules[lookup]))
	}
	return append(res, s[len(s)-1])
}

func (m *Memoizer) Lookup(a, b rune) rune {
	return m.Rules[string(a)+string(b)]
}
