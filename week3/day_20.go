package week3

import "fmt"

// TrenchMap 4924 too high
// Answer is 4917
func TrenchMap(f string) int {
	trench := NewImageEnhancer(f)
	trench.Enhance()
	trench.Enhance()
	fmt.Println(trench)

	return 0
}
