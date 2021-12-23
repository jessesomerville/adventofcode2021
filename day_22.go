package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	//go:embed inputs/day_22.txt
	reactorFile string

	//go:embed test_inputs/day_22.txt
	reactorFileTest string
)

func reactorRebootFull() int {
	// steps := parseReactorSteps(reactorFileTest).Steps
	return 0
}

func reactorReboot() int {
	reactorDim := 101
	steps := parseReactorSteps(reactorFileTest)
	reactorGrid := make([]int, int(math.Pow(float64(reactorDim), 3)))

	onCount := 0
	for _, step := range steps.Steps {
		if !step.Init {
			continue
		}
		for x := step.C.X1; x <= step.C.X2; x++ {
			for y := step.C.Y1; y <= step.C.Y2; y++ {
				for z := step.C.Z1; z <= step.C.Z2; z++ {
					coord := coords(x, y, z, reactorDim)
					if step.On {
						if reactorGrid[coord] == 0 {
							reactorGrid[coord] = 1
							onCount++
						}
					} else if reactorGrid[coord] == 1 {
						reactorGrid[coord] = 0
						onCount--
					}
				}
			}
		}
	}
	return onCount
}

func coords(x, y, z int, width int) int {
	return (z * width * width) + (y * width) + x
}

// Reboot represents the steps to reboot the reactor.
type Reboot struct {
	Steps []*Step
}

// Step represents a single reactor reboot step.
type Step struct {
	On   bool
	C    *Cuboid
	Init bool
}

func (s *Step) String() string {
	buf := new(strings.Builder)
	if s.On {
		fmt.Fprint(buf, "on  ")
	} else {
		fmt.Fprint(buf, "off ")
	}

	fmt.Fprintf(buf, "x=%d..%d,", s.C.X1, s.C.X2)
	fmt.Fprintf(buf, "y=%d..%d,", s.C.Y1, s.C.Y2)
	fmt.Fprintf(buf, "z=%d..%d\n", s.C.Z1, s.C.Z2)
	return buf.String()
}

func (s *Step) Initializer() {
	if inMin(s.C.X1, s.C.Y1, s.C.Z1) && inMax(s.C.X2, s.C.Y2, s.C.Z2) {
		s.Init = true
	}
}

func inMin(nums ...int) bool {
	for _, i := range nums {
		if i < 0 {
			return false
		}
	}
	return true
}

func inMax(nums ...int) bool {
	for _, i := range nums {
		if i >= 101 {
			return false
		}
	}
	return true
}

func parseReactorSteps(f string) *Reboot {
	lines := strings.Split(f, "\n")
	steps := make([]*Step, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		instruction := parts[0]
		ranges := strings.Split(parts[1], ",")
		xmin, xmax := intsFromStr(ranges[0])
		ymin, ymax := intsFromStr(ranges[1])
		zmin, zmax := intsFromStr(ranges[2])
		step := &Step{
			On: instruction == "on",
			C: &Cuboid{
				X1: xmin,
				X2: xmax,
				Y1: ymin,
				Y2: ymax,
				Z1: zmin,
				Z2: zmax,
			},
		}
		step.Initializer()
		steps = append(steps, step)
	}
	return &Reboot{
		Steps: steps,
	}
}

func intsFromStr(s string) (int, int) {
	rangeStr := s[2:]
	numsStr := strings.Split(rangeStr, "..")
	min, _ := strconv.Atoi(numsStr[0])
	max, _ := strconv.Atoi(numsStr[1])
	return min + 50, max + 50
}
