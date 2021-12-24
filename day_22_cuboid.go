package main

import (
	"fmt"
	"math"
	"strings"
)

// Cuboid represents a cuboid.
type Cuboid struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
}

// Clip clips c at its intersection with cr (clipping-reference) and returns
// the result of the dissection of the resulting rectilinear shape into cuboids.
func (c *Cuboid) Clip(cr *Cuboid) []*Cuboid {
	if !c.Intersect(cr) {
		return []*Cuboid{c}
	}
	splits := make([]*Cuboid, 0, 27)
	xSplits := c.SplitX(cr)
	for _, sx := range xSplits {
		if !sx.Intersect(cr) {
			splits = append(splits, sx)
			continue
		}
		ySplits := sx.SplitY(cr)
		for _, sy := range ySplits {
			if !sy.Intersect(cr) {
				splits = append(splits, sy)
				continue
			}
			zSplits := sy.SplitZ(cr)
			for _, sz := range zSplits {
				if !sz.Intersect(cr) {
					splits = append(splits, sz)
				}
			}
		}
	}

	return splits
}

func (c Cuboid) Volume() int {
	xDist := math.Abs(float64(c.X2-c.X1)) + 1
	yDist := math.Abs(float64(c.Y2-c.Y1)) + 1
	zDist := math.Abs(float64(c.Z2-c.Z1)) + 1

	return int(xDist * yDist * zDist)
}

type Bounds struct {
	Min, Max int
}

func (c Cuboid) SplitX(cr *Cuboid) []*Cuboid {
	if (cr.X1 == c.X1 && cr.X2 == c.X2) || (c.X1 >= cr.X1 && c.X2 <= cr.X2) {
		return []*Cuboid{&c}
	}
	if (c.X1 == cr.X1 && c.X2 < cr.X2) || (c.X1 > cr.X1 && c.X2 == cr.X2) {
		return []*Cuboid{&c}
	}

	xBounds := getBounds(c.X1, c.X2, cr.X1, cr.X2)
	spl := make([]*Cuboid, 0, 3)
	for _, b := range xBounds {
		subCuboid := &Cuboid{
			b.Min, b.Max,
			c.Y1, c.Y2,
			c.Z1, c.Z2,
		}
		spl = append(spl, subCuboid)
	}

	return spl
}

func (c Cuboid) SplitY(cr *Cuboid) []*Cuboid {
	if (cr.Y1 == c.Y1 && cr.Y2 == c.Y2) || (c.Y1 >= cr.Y1 && c.Y2 <= cr.Y2) {
		return []*Cuboid{&c}
	}
	if (c.Y1 == cr.Y1 && c.Y2 < cr.Y2) || (c.Y1 > cr.Y1 && c.Y2 == cr.Y2) {
		return []*Cuboid{&c}
	}

	yBounds := getBounds(c.Y1, c.Y2, cr.Y1, cr.Y2)
	spl := make([]*Cuboid, 0, 3)
	for _, b := range yBounds {
		subCuboid := &Cuboid{
			c.X1, c.X2,
			b.Min, b.Max,
			c.Z1, c.Z2,
		}
		spl = append(spl, subCuboid)
	}

	return spl
}

func (c Cuboid) SplitZ(cr *Cuboid) []*Cuboid {
	if (cr.Z1 == c.Z1 && cr.Z2 == c.Z2) || (c.Z1 >= cr.Z1 && c.Z2 <= cr.Z2) {
		return []*Cuboid{&c}
	}
	if (c.Z1 == cr.Z1 && c.Z2 < cr.Z2) || (c.Z1 > cr.Z1 && c.Z2 == cr.Z2) {
		return []*Cuboid{&c}
	}

	zBounds := getBounds(c.Z1, c.Z2, cr.Z1, cr.Z2)
	spl := make([]*Cuboid, 0, 3)
	for _, b := range zBounds {
		subCuboid := &Cuboid{
			c.X1, c.X2,
			c.Y1, c.Y2,
			b.Min, b.Max,
		}
		spl = append(spl, subCuboid)
	}

	return spl
}

// Intersect performs collision detection between two cuboids.
func (c0 *Cuboid) Intersect(c1 *Cuboid) bool {
	xInt := anyIn(c0.X1, c0.X2, c1.X1, c1.X2)
	yInt := anyIn(c0.Y1, c0.Y2, c1.Y1, c1.Y2)
	zInt := anyIn(c0.Z1, c0.Z2, c1.Z1, c1.Z2)

	return xInt && yInt && zInt

	// x := max(min(c0.X2, c1.X2)-max(c0.X1, c1.X1), 0)
	// y := max(min(c0.Y2, c1.Y2)-max(c0.Y1, c1.Y1), 0)
	// z := max(min(c0.Z2, c1.Z2)-max(c0.Z1, c1.Z1), 0)
	// return (x * y * z) > 0

	// xInt := !(c0.X2 < c1.X1) && !(c1.X2 < c0.X1)
	// yInt := !(c0.Y2 < c1.Y1) && !(c1.Y2 < c0.Y1)
	// zInt := !(c0.Z2 < c1.Z1) && !(c1.Z2 < c0.Z1)
	// return xInt || yInt || zInt
}

func anyIn(a0, a1, b0, b1 int) bool {
	aDist := math.Abs(float64(a1 - a0))
	bDist := math.Abs(float64(b1 - b0))

	if bDist < aDist {
		a0, a1, b0, b1 = b0, b1, a0, a1
	}

	for i := a0; i <= a1; i++ {
		if i == b0 || i == b1 {
			return true
		}
	}
	return false
}

func (c *Cuboid) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "X: %d..%d\n", c.X1, c.X2)
	fmt.Fprintf(buf, "Y: %d..%d\n", c.Y1, c.Y2)
	fmt.Fprintf(buf, "Z: %d..%d\n", c.Z1, c.Z2)
	return buf.String()
}

func getBounds(a0, a1, b0, b1 int) []*Bounds {
	if a0 < b0 && a1 > b1 {
		// a complete encompasses b
		return []*Bounds{
			{a0, b0 - 1},
			{b0, b1},
			{b1 + 1, a1},
		}
	}
	if a0 < b0 {
		// a0 | b0 | a1 | b1
		return []*Bounds{
			{a0, b0 - 1},
			{b0, a1},
		}
	} else {
		// b0 | a0 | b1 | a1
		return []*Bounds{
			{a0, b1},
			{b1 + 1, a1},
		}
	}
}
