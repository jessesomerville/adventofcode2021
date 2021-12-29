package week4

type InstBlock interface {
	Exec(int, int) int
	AllWs(int) []int
}

type AddInst struct {
	A1, A2 int
}

func (pi *AddInst) Exec(w, z int) int {
	if (z%26)+pi.A1 != w {
		z = (z * 26) + w + pi.A2
	}
	return z
}

func (pi *AddInst) AllWs(z int) []int {
	out := make([]int, 0, 9)
	for w := 1; w <= 9; w++ {
		out = append(out, pi.Exec(w, z))
	}
	return out
}

type SubInst struct {
	A1, A2 int
}

func (mi *SubInst) Exec(w, z int) int {
	x := z
	z /= 26
	if (x%26)-mi.A1 != w {
		z = (z * 26) + w + mi.A2
	}
	return z
}

func (si *SubInst) AllWs(z int) []int {
	out := make([]int, 0, 9)
	for w := 1; w <= 9; w++ {
		out = append(out, si.Exec(w, z))
	}
	return out
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
