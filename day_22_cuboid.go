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

func (c Cuboid) SplitX(cr *Cuboid) []*Cuboid {
	if (c.X1 == cr.X1) && (c.X2 == cr.X2) {
		return []*Cuboid{&c}
	}
	xInt0, xInt1 := max(c.X1, cr.X1), min(c.X2, cr.X2)
	xBounds := SortUniq(xInt0, xInt1, c.X1, c.X2)
	spl := make([]*Cuboid, 0, 3)
	for xi := 0; xi < len(xBounds)-1; xi++ {
		subCuboid := &Cuboid{
			xBounds[xi] + 1, xBounds[xi+1],
			c.Y1, c.Y2,
			c.Z1, c.Z2,
		}
		if xi == 0 {
			subCuboid.X1--
		}
		spl = append(spl, subCuboid)
	}
	return spl
}

func (c Cuboid) SplitY(cr *Cuboid) []*Cuboid {
	if (c.Y1 == cr.Y1) && (c.Y2 == cr.Y2) {
		return []*Cuboid{&c}
	}
	y0, y1 := max(c.Y1, cr.Y1), min(c.Y2, cr.Y2)
	yBounds := SortUniq(y0, y1, c.Y1, c.Y2)
	spl := make([]*Cuboid, 0, 3)
	for i := 0; i < len(yBounds)-1; i++ {
		subCuboid := &Cuboid{
			c.X1, c.X2,
			yBounds[i] + 1, yBounds[i+1],
			c.Z1, c.Z2,
		}
		if i == 0 {
			subCuboid.Y1--
		}
		spl = append(spl, subCuboid)
	}
	return spl
}

func (c Cuboid) SplitZ(cr *Cuboid) []*Cuboid {
	if (c.Z1 == cr.Z1) && (c.Z2 == cr.Z2) {
		return []*Cuboid{&c}
	}
	z0, z1 := max(c.Z1, cr.Z1), min(c.Z2, cr.Z2)
	zBounds := SortUniq(z0, z1, c.Z1, c.Z2)
	spl := make([]*Cuboid, 0, 3)
	for i := 0; i < len(zBounds)-1; i++ {
		subCuboid := &Cuboid{
			c.X1, c.X2,
			c.Y1, c.Y2,
			zBounds[i] + 1, zBounds[i+1],
		}
		if i == 0 {
			subCuboid.Z1--
		}
		spl = append(spl, subCuboid)
	}
	return spl
}

// Intersect performs collision detection between two cuboids.
func (c0 *Cuboid) Intersect(c1 *Cuboid) bool {
	x := max(min(c0.X2, c1.X2)-max(c0.X1, c1.X1), 0)
	y := max(min(c0.Y2, c1.Y2)-max(c0.Y1, c1.Y1), 0)
	z := max(min(c0.Z2, c1.Z2)-max(c0.Z1, c1.Z1), 0)
	return (x * y * z) > 0
}

func (c *Cuboid) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "X: %d..%d\n", c.X1, c.X2)
	fmt.Fprintf(buf, "Y: %d..%d\n", c.Y1, c.Y2)
	fmt.Fprintf(buf, "Z: %d..%d\n", c.Z1, c.Z2)
	return buf.String()
}
