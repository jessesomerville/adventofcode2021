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
	steps := parseReactorSteps(reactorFile)
	var newSteps []*Step
	for _, c1 := range steps {
		for _, c2 := range newSteps {
			x1, x0 := min(c1.X2, c2.X2), max(c1.X1, c2.X1)
			y1, y0 := min(c1.Y2, c2.Y2), max(c1.Y1, c2.Y1)
			z1, z0 := min(c1.Z2, c2.Z2), max(c1.Z1, c2.Z1)
			if x0 < x1 && y0 < y1 && z0 < z1 {
				newSteps = append(newSteps, &Step{&Cuboid{x0, x1, y0, y1, z0, z1}, !c2.On})
			}
		}
		if c1.On {
			newSteps = append(newSteps, c1)
		}
	}
	lights := 0
	for _, s := range newSteps {
		sign := 1
		if !s.On {
			sign = -1
		}
		lights += sign * s.GetVolume()
	}
	fmt.Println(lights)

	return 0
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
			NewCuboid(xmin, xmax, ymin, ymax, zmin, zmax),
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

// Cuboid represents a cuboid.
type Cuboid struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
}

func NewCuboid(x1, x2, y1, y2, z1, z2 int) *Cuboid {
	c := &Cuboid{x1, x2, y1, y2, z1, z2}
	return c
}

func (c1 *Cuboid) GetVolume() int {
	x := math.Abs(float64(c1.X2-c1.X1)) + 1
	y := math.Abs(float64(c1.Y2-c1.Y1)) + 1
	z := math.Abs(float64(c1.Z2-c1.Z1)) + 1
	return int(x * y * z)
}
