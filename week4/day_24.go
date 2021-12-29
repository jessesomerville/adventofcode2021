package week4

import (
	"fmt"
)

// MONAD finds the largest valid MONAD accepted by the ALU.
func MONAD(f string) {
	validMONADS := ReverseMONAD(len(ops)-1, 0)
	maxMONAD := 0
	minMONAD := int(^uint(0) >> 1)
	for _, m := range validMONADS {
		if m.MONAD > maxMONAD {
			maxMONAD = m.MONAD
		}
		if m.MONAD < minMONAD {
			minMONAD = m.MONAD
		}
	}
	fmt.Println("MAX:", maxMONAD, "MIN:", minMONAD)
}

// ReverseMONAD recursively finds valid MONADs by running the program in reverse.
func ReverseMONAD(opIdx, targetOut int) []*ProgramInput {
	if opIdx < 0 {
		return nil
	}
	out := []*ProgramInput{}

	op := ops[opIdx]
	inputs := op.GenInput(targetOut)
	for w, inp := range inputs {
		if opIdx == 0 {
			out = append(out, &ProgramInput{w, inp})
		} else {
			for _, prev := range ReverseMONAD(opIdx-1, inp) {
				prev.MONAD = (prev.MONAD * 10) + w
				out = append(out, prev)
			}
		}
	}

	return out
}

// ProgramInput represents the MONAD input and the initial z value to use when reversing the MONAD.
type ProgramInput struct {
	MONAD, Z int
}
