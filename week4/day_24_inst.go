package week4

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Register struct {
	Value int
}

type ALU struct {
	W, X, Y, Z *Register
}

func (alu *ALU) ValidMONAD() bool {
	return alu.Z.Value == 0
}

func (alu *ALU) Execute(program string, input int) {
	if invalidInput(input) {
		return
	}
	instructions := strings.Split(program, "\n")

	for _, inst := range instructions {
		cmd := strings.Split(inst, " ")
		switch cmd[0] {
		case "inp":
			// fmt.Println(alu.Z.Value)
			var msd int
			msd, input = splitMSD(input)
			if msd == 0 {
				log.Fatal("input contained a 0 or was too short for program")
			}
			alu.Inp(cmd[1], msd)
		case "add":
			alu.Add(cmd[1], cmd[2])
		case "mul":
			alu.Mul(cmd[1], cmd[2])
		case "div":
			alu.Div(cmd[1], cmd[2])
		case "mod":
			alu.Mod(cmd[1], cmd[2])
		case "eql":
			alu.Eql(cmd[1], cmd[2])
		default:
			log.Fatalf("Unknown instruction: %q", cmd[0])
		}
	}
}

func (alu *ALU) Inp(a string, val int) {
	alu.RegisterByName(a).Value = val
}

func (alu *ALU) Add(a, b string) {
	alu.RegisterByName(a).Value += alu.Value(b)
}

func (alu *ALU) Mul(a, b string) {
	alu.RegisterByName(a).Value *= alu.Value(b)
}

func (alu *ALU) Div(a, b string) {
	bVal := alu.Value(b)
	if bVal == 0 {
		log.Fatalf("attempted to divide %s by 0", a)
	}
	alu.RegisterByName(a).Value /= bVal
}

func (alu *ALU) Mod(a, b string) {
	aReg := alu.RegisterByName(a)
	bVal := alu.Value(b)
	if aReg.Value < 0 || bVal <= 0 {
		log.Fatalf("attempted mod with invalid params: (%d %% %d)", aReg.Value, bVal)
	}
	aReg.Value %= bVal
}

func (alu *ALU) Eql(a, b string) {
	aReg := alu.RegisterByName(a)
	if aReg.Value == alu.Value(b) {
		aReg.Value = 1
	} else {
		aReg.Value = 0
	}
}

func (alu *ALU) RegisterByName(name string) *Register {
	switch name {
	case "w":
		return alu.W
	case "x":
		return alu.X
	case "y":
		return alu.Y
	case "z":
		return alu.Z
	default:
		log.Fatalf("ALU was passed an unknown variable: %q", name)
	}
	return nil
}

func (alu *ALU) Value(a string) int {
	if val, err := strconv.Atoi(a); err == nil {
		return val
	}
	return alu.RegisterByName(a).Value
}

func NewALU() *ALU {
	return &ALU{
		new(Register),
		new(Register),
		new(Register),
		new(Register),
	}
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

func (alu *ALU) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "W: %d\n", alu.W.Value)
	fmt.Fprintf(buf, "X: %d\n", alu.X.Value)
	fmt.Fprintf(buf, "Y: %d\n", alu.Y.Value)
	fmt.Fprintf(buf, "Z: %d\n", alu.Z.Value)
	return buf.String()
}

func (alu *ALU) Reset() {
	alu.W.Value = 0
	alu.X.Value = 0
	alu.Y.Value = 0
	alu.Z.Value = 0
}

func invalidInput(x int) bool {
	return strings.ContainsRune(fmt.Sprint(x), '0')
}
