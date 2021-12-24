package main

import "fmt"

// Point3D represents the coordinates of a point in 3D space.
type Point3D struct {
	X, Y, Z int
}

func (p *Point3D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}
