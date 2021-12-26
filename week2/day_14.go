package week2

import (
	"fmt"
	"strings"
	"sync"
)

func Polymerization(f string) {
	after5 := step20(f)

	parts := strings.Split(f, "\n\n")
	rules := rulesMap(parts[1])

	ruleMapExt := &RuleMap{
		sync.RWMutex{},
		make(map[string]map[rune]int, len(rules)),
	}
	var wg sync.WaitGroup
	for rule := range rules {
		wg.Add(1)
		go func(rule string) {
			mem := NewMemoizer(rulesMap(parts[1]))
			res := rule
			for i := 0; i < 20; i++ {
				res = string(mem.Evaluate([]rune(res)))
			}
			res = res[1:]
			rc := make(map[rune]int)
			for _, r := range res {
				rc[r]++
			}
			ruleMapExt.Write(rule, rc)
			wg.Done()
		}(rule)
	}
	wg.Wait()

	runeCounts := make(map[rune]int)
	for i := 0; i < len(after5)-1; i++ {
		cnts := ruleMapExt.rules[after5[i:i+2]]
		for r, count := range cnts {
			runeCounts[r] += count
		}
	}

	min := int(^uint(0) >> 1)
	max := 0
	for _, c := range runeCounts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	fmt.Println(max - min)
}

func step20(f string) string {
	parts := strings.Split(f, "\n\n")
	polyTmpl := []rune(parts[0])
	rules := rulesMap(parts[1])

	mem := NewMemoizer(rules)
	for i := 0; i < 20; i++ {
		polyTmpl = mem.Evaluate(polyTmpl)
	}
	return string(polyTmpl)
}

type RuleMap struct {
	sync.RWMutex
	rules map[string]map[rune]int
}

func (rm *RuleMap) Write(key string, val map[rune]int) {
	rm.Lock()
	defer rm.Unlock()
	rm.rules[key] = val
}

func rulesMap(s string) map[string]rune {
	lines := strings.Split(s, "\n")
	ruleMap := make(map[string]rune, len(lines))
	for _, line := range lines {
		spl := strings.Split(line, " -> ")
		ruleMap[spl[0]] = rune(spl[1][0])
	}
	return ruleMap
}
