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
	steps := parseReactorSteps(reactorFileTest)
	sections := sectionByOff(steps)

	for _, sec := range sections {
		for _, i := range sec {
			fmt.Println(i)
		}
		fmt.Println("------------")
	}

	// var onSteps []*Cuboid
	// for _, section := range sections {
	// 	onSteps = append(onSteps, removeOff(section)...)
	// }
	// var reduced []*Cuboid
	// for i := 0; i < len(onSteps)-1; i++ {
	// 	r := reduce(onSteps[i], onSteps[i+1:])
	// 	reduced = append(reduced, r...)
	// }
	// reduced = append(reduced, onSteps[len(onSteps)-1])
	// lights := 0
	// for _, r := range reduced {
	// 	// fmt.Println(r)
	// 	lights += r.Volume()
	// }
	// fmt.Println(lights)

	return 0
}

func reduce(s0 *Cuboid, steps []*Cuboid) []*Cuboid {
	if len(steps) == 0 {
		return []*Cuboid{s0}
	}
	var reduced []*Cuboid
	for _, c := range s0.Clip(steps[0]) {
		reduced = append(reduced, reduce(c, steps[1:])...)
	}
	return reduced
}

func removeOff(section []*Step) []*Cuboid {
	if section[len(section)-1].On {
		return stepsToCuboids(section)
	}

	idx := 0
	for {
		if idx == len(section)-1 {
			break
		}
		if !section[idx].On {
			break
		}
		idx++
	}
	offs := stepsToCuboids(section[idx:])
	ons := stepsToCuboids(section[:idx])
	for _, off := range offs {
		ons = turnOff(off, ons)
	}
	return ons
}

func turnOff(off *Cuboid, ons []*Cuboid) []*Cuboid {
	var turnedOff []*Cuboid
	for _, on := range ons {
		turnedOff = append(turnedOff, on.Clip(off)...)
	}
	return turnedOff
}

func stepsToCuboids(s []*Step) []*Cuboid {
	c := make([]*Cuboid, 0, len(s))
	for _, step := range s {
		c = append(c, step.Cuboid)
	}
	return c
}

func sectionByOff(steps []*Step) [][]*Step {
	var sections [][]*Step
	var section []*Step
	for i, s := range steps {
		section = append(section, s)
		if !s.On {
			for j := i + 1; j < len(steps); j++ {
				if !steps[j].On {
					section = append(section, steps[j])
				}
			}
			sections = append(sections, section)
			section = []*Step{}
		}
	}
	if len(section) > 0 {
		sections = append(sections, section)
	}
	return sections
}

func reactorReboot() int {
	reactorDim := 101
	steps := parseReactorSteps(reactorFileTest)
	reactorGrid := make([]int, int(math.Pow(float64(reactorDim), 3)))

	onCount := 0
	for _, step := range steps {
		if step.X1 < -50 || step.Y1 < -50 || step.Z1 < -50 || step.X2 > 50 || step.Y2 > 50 || step.Z2 > 50 {
			continue
		}
		for x := step.X1; x <= step.X2; x++ {
			for y := step.Y1; y <= step.Y2; y++ {
				for z := step.Z1; z <= step.Z2; z++ {
					coord := coords(x+50, y+50, z+50, reactorDim)
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

// Step represents a single reactor reboot step.
type Step struct {
	*Cuboid
	On bool
}

func parseReactorSteps(f string) []*Step {
	lines := strings.Split(f, "\n")
	steps := make([]*Step, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		ranges := strings.Split(parts[1], ",")
		xmin, xmax := intsFromStr(ranges[0])
		ymin, ymax := intsFromStr(ranges[1])
		zmin, zmax := intsFromStr(ranges[2])
		step := &Step{
			&Cuboid{
				X1: xmin,
				X2: xmax,
				Y1: ymin,
				Y2: ymax,
				Z1: zmin,
				Z2: zmax,
			},
			parts[0] == "on",
		}
		steps = append(steps, step)
	}
	return steps
}

func intsFromStr(s string) (int, int) {
	rangeStr := s[2:]
	numsStr := strings.Split(rangeStr, "..")
	min, _ := strconv.Atoi(numsStr[0])
	max, _ := strconv.Atoi(numsStr[1])
	return min, max
}
