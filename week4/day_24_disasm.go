package week4

import "math"

// InstBlock represents an instruction block in the program.
type InstBlock interface {
	Exec(int, int) int
	GenInput(int) map[int]int
}

// AddInst represents an instruction block that performs an addition check.
type AddInst struct {
	A1, A2 int
}

// Exec executes the addition instruction block.
func (pi *AddInst) Exec(w, z int) int {
	if (z%26)+pi.A1 != w {
		z = (z * 26) + w + pi.A2
	}
	return z
}

// GenInput generates the inputs for the given target.
func (pi *AddInst) GenInput(target int) map[int]int {
	inputs := make(map[int]int, 9)
	for w := 1; w <= 9; w++ {
		pre := target - w - pi.A2
		if pre%26 == 0 {
			inputs[w] = pre / 26
		}
	}
	return inputs
}

// SubInst represents an instruction block that performs a subtraction check.
type SubInst struct {
	A1, A2 int
}

// Exec executes the subtraction instruction block.
func (si *SubInst) Exec(w, z int) int {
	x := z
	z /= 26
	if (x%26)-si.A1 != w {
		z = (z * 26) + w + si.A2
	}
	return z
}

// GenInput generates the inputs for the given target.
func (si *SubInst) GenInput(target int) map[int]int {
	maxInput := (target + 1) * 26
	inputs := make(map[int]int, 9)
	for w := 1; w <= 9; w++ {
		modTarget := 26 - (si.A1 + w)
		inputs[w] = maxInput - modTarget
	}
	return inputs
}

var ops = []InstBlock{
	&AddInst{13, 8},
	&AddInst{12, 16},
	&AddInst{10, 4},
	&SubInst{11, 1},
	&AddInst{14, 13},
	&AddInst{13, 5},
	&AddInst{12, 0},
	&SubInst{5, 10},
	&AddInst{10, 7},
	&SubInst{0, 2},
	&SubInst{11, 13},
	&SubInst{13, 15},
	&SubInst{13, 14},
	&SubInst{11, 9},
}

// ExecuteProgram runs the instruction blocks consecutively.
func ExecuteProgram(monad, z int) int {
	var w int
	for _, op := range ops {
		w, monad = splitMSD(monad)
		z = op.Exec(w, z)
	}
	return z
}

func splitMSD(input int) (int, int) {
	x := float64(input)
	lg := int(math.Pow10(int(math.Log10(x))))
	if lg == 0 {
		return 0, 0
		// log.Fatalf("input contained a zero: %d", input)
	}
	return input / lg, input % lg
}
