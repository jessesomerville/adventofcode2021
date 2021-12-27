package week3

import "fmt"

// TrenchMap is the solution for day 20.
func TrenchMap(f string) int {
	trench := NewImageEnhancer(f)
	trench.Scale(100)
	for i := 0; i < 50; i++ {
		trench.Enhance()
	}
	trench.Trim(50)
	fmt.Println(trench.PixelCount())

	return 0
}
