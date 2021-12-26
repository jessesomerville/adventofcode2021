package week4

import "fmt"

func MONAD(f string) {
	op := ops[len(ops)-2]
	for in := 21*26 - 1; in >= 337; in-- {
		for w := 1; w <= 9; w++ {
			res := op.Op(w, in)
			if res >= 12 && res <= 20 {
				fmt.Println(w, in)
			}
		}
	}

	// z := 9541593
	// for _, op := range ops {
	// 	fmt.Println(z)
	// 	z = op.Op(9, z)
	// }
	// fmt.Println(z)

	// Loop:
	// 	for i := 999999999; i >= 988888888; i-- {
	// 		if i%10 == 0 {
	// 			continue
	// 		}
	// 		input := i
	// 		var z int
	// 		var val int
	// 		for j := 0; j < 9; j++ {
	// 			val, input = splitMSD(input)
	// 			if val == 0 {
	// 				continue Loop
	// 			}
	// 			z = ops[j].Op(val, z)
	// 		}
	// 		if z == 9541593 {
	// 			fmt.Println(i)
	// 			return
	// 		}
	// 	}

	// input := 111111111
	// var z, val int
	// for i := 0; i < 9; i++ {
	// 	val, input = splitMSD(input)
	// 	z = ops[i].Op(val, z)
	// }
	// fmt.Println(z)

}

type InstBlock interface {
	Op(int, int) int
}

type AddInst struct {
	A1, A2 int
}

func (pi *AddInst) Op(w, z int) int {
	if (z%26)+pi.A1 != w {
		z = (z * 26) + w + pi.A2
	}
	return z
}

type SubInst struct {
	A1, A2 int
}

func (mi *SubInst) Op(w, z int) int {
	x := z
	z /= 26
	if (x%26)-mi.A1 != w {
		z = (z * 26) + w + mi.A2
	}
	return z
}

var ops = []InstBlock{
	// &AddInst{13, 8},
	// &AddInst{12, 16},
	// &AddInst{10, 4},
	// &SubInst{11, 1},
	// &AddInst{14, 13},
	// &AddInst{13, 5},
	// &AddInst{12, 0},
	// &SubInst{5, 10},
	// &AddInst{10, 7},
	&SubInst{0, 2},
	&SubInst{11, 13},
	&SubInst{13, 15},
	&SubInst{13, 14},
	&SubInst{11, 9},
}
