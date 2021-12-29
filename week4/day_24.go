package week4

import (
	"fmt"
	"math"
)

func MONAD(f string) {

	// targets := map[int]bool{0: true}
	// maxTarget := 1
	// for opIdx := len(ops) - 1; opIdx >= 0; opIdx-- {
	// 	targets, maxTarget = findBounds(ops[opIdx], targets, maxTarget)
	// }

	wz := WZ{123, 456}
	nwz := wz.Update(9, 678)
	fmt.Println(nwz.W, nwz.Z)

	// input := 111111111
	// var z, val int
	// for i := 0; i < 9; i++ {
	// 	val, input = splitMSD(input)
	// 	z = ops[i].Op(val, z)
	// }
	// fmt.Println(z)

}

type MONADRev struct {
	Ops         []InstBlock
	ValidMONADs []*WZ
}

func (mr *MONADRev) Reverse(opIdx, targetOut int) []*WZ {
	// op := mr.Ops[opIdx]
	// for z := 0; z <= (targetOut+1)*26; z++ {
	// 	for w, res := range op.AllWs(z) {
	// 		if res == targetOut {

	// 		}
	// 	}
	// }
	return nil
}

type WZ struct {
	W, Z int
}

func (wz WZ) Update(w, z int) *WZ {
	nextPow := int(math.Log10(float64(wz.W))) + 1
	newW := (wz.W) + (w * int(math.Pow10(nextPow)))
	return &WZ{
		W: newW,
		Z: z,
	}
}

func findBounds(op InstBlock, targets map[int]bool, maxTarget int) (map[int]bool, int) {
	switch inst := op.(type) {
	case *AddInst:
		op = inst
	case *SubInst:
		op = inst
	}

	allowedInputs := make(map[int]bool)
	maxInput := -1
	for z := 0; z <= (maxTarget+1)*26; z++ {
		for _, res := range op.AllWs(z) {
			if targets[res] {
				allowedInputs[z] = true
				if z > maxInput {
					maxInput = z
				}
			}
		}
	}
	return allowedInputs, maxInput
}
