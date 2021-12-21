package main

import (
	"math"
)

const (
	Pos0   = 0
	Pos90  = 1.5707963267948966
	Pos180 = 3.141592653589793
	Pos270 = 4.71238898038469
)

var (
	Rx0 = [][]int{
		{1, 0, 0},
		{0, int(math.Cos(Pos0)), -int(math.Sin(Pos0))},
		{0, int(math.Sin(Pos0)), int(math.Cos(Pos0))},
	}
	Rx90 = [][]int{
		{1, 0, 0},
		{0, int(math.Cos(Pos90)), -int(math.Sin(Pos90))},
		{0, int(math.Sin(Pos90)), int(math.Cos(Pos90))},
	}
	Rx180 = [][]int{
		{1, 0, 0},
		{0, int(math.Cos(Pos180)), -int(math.Sin(Pos180))},
		{0, int(math.Sin(Pos180)), int(math.Cos(Pos180))},
	}
	Rx270 = [][]int{
		{1, 0, 0},
		{0, int(math.Cos(Pos270)), -int(math.Sin(Pos270))},
		{0, int(math.Sin(Pos270)), int(math.Cos(Pos270))},
	}

	Ry0 = [][]int{
		{int(math.Cos(Pos0)), 0, int(math.Sin(Pos0))},
		{0, 1, 0},
		{-int(math.Sin(Pos0)), 0, int(math.Cos(Pos0))},
	}
	Ry90 = [][]int{
		{int(math.Cos(Pos90)), 0, int(math.Sin(Pos90))},
		{0, 1, 0},
		{-int(math.Sin(Pos90)), 0, int(math.Cos(Pos90))},
	}
	Ry180 = [][]int{
		{int(math.Cos(Pos180)), 0, int(math.Sin(Pos180))},
		{0, 1, 0},
		{-int(math.Sin(Pos180)), 0, int(math.Cos(Pos180))},
	}
	Ry270 = [][]int{
		{int(math.Cos(Pos270)), 0, int(math.Sin(Pos270))},
		{0, 1, 0},
		{-int(math.Sin(Pos270)), 0, int(math.Cos(Pos270))},
	}

	Rz0 = [][]int{
		{int(math.Cos(Pos0)), -int(math.Sin(Pos0)), 0},
		{int(math.Sin(Pos0)), int(math.Cos(Pos0)), 0},
		{0, 0, 1},
	}
	Rz90 = [][]int{
		{int(math.Cos(Pos90)), -int(math.Sin(Pos90)), 0},
		{int(math.Sin(Pos90)), int(math.Cos(Pos90)), 0},
		{0, 0, 1},
	}
	Rz180 = [][]int{
		{int(math.Cos(Pos180)), -int(math.Sin(Pos180)), 0},
		{int(math.Sin(Pos180)), int(math.Cos(Pos180)), 0},
		{0, 0, 1},
	}
	Rz270 = [][]int{
		{int(math.Cos(Pos270)), -int(math.Sin(Pos270)), 0},
		{int(math.Sin(Pos270)), int(math.Cos(Pos270)), 0},
		{0, 0, 1},
	}
)

func rotateMatrix(origCoord []int, b1 *beacon) (x, y, z int) {
	coordToMatch := []int{b1.X, b1.Y, b1.Z}
	return checkAllRotations(origCoord, coordToMatch)
}

func checkAllRotations(c0, c1 []int) (xr, yr, zr int) {
	rotXMatrices := [][][]int{Rx0, Rx90, Rx180, Rx270}
	rotYMatrices := [][][]int{Ry0, Ry90, Ry180, Ry270}
	rotZMatrices := [][][]int{Rz0, Rz90, Rz180, Rz270}

	for xr, xri := range rotXMatrices {
		for yr, yri := range rotYMatrices {
			for zr, zri := range rotZMatrices {
				newCoords := rotate(c0, xri)
				newCoords = rotate(newCoords, yri)
				newCoords = rotate(newCoords, zri)
				if correctRotation(newCoords, c1) {
					return xr, yr, zr
				}
			}
		}
	}
	return -1, -1, -1
}

func correctRotation(c0, c1 []int) bool {
	x1, y1, z1 := c0[0], c0[1], c0[2]
	x2, y2, z2 := c1[0], c1[1], c1[2]
	xDiff := x1 == x2
	yDiff := y1 == y2
	zDiff := z1 == z2
	if xDiff && yDiff && zDiff {
		return true
	}
	return false
}

// func rotateMatrix2(c0, c1 []int, b0, b1 *beacon) (x, y, z int) {
// 	b0c := []int{b0.X, b0.Y, b0.Z}
// 	b1c := []int{b1.X, b1.Y, b1.Z}
// 	return checkAllRotations2(c0, c1, b0c, b1c)
// }

// func checkAllRotations2(c0, c1, b0, b1 []int) (xr, yr, zr int) {
// 	rotXMatrices := [][][]int{Rx0, Rx90, Rx180, Rx270}
// 	rotYMatrices := [][][]int{Ry0, Ry90, Ry180, Ry270}
// 	rotZMatrices := [][][]int{Rz0, Rz90, Rz180, Rz270}

// 	for xr, xri := range rotXMatrices {
// 		for yr, yri := range rotYMatrices {
// 			for zr, zri := range rotZMatrices {
// 				newCoords0 := rotate(c0, xri)
// 				newCoords0 = rotate(newCoords0, yri)
// 				newCoords0 = rotate(newCoords0, zri)
// 				newCoords1 := rotate(c1, xri)
// 				newCoords1 = rotate(newCoords1, yri)
// 				newCoords1 = rotate(newCoords1, zri)
// 				if correctRotation2(newCoords0, newCoords1, b0, b1) {
// 					return xr, yr, zr
// 				}
// 			}
// 		}
// 	}
// 	return -1, -1, -1
// }

// func correctRotation2(c0, c1, b0, b1 []int) bool {
// 	x0, y0, z0 := c0[0], c0[1], c0[2]
// 	x1, y1, z1 := c1[0], c1[1], c1[2]
// 	xb0, yb0, zb0 := b0[0], b0[1], b0[2]
// 	xb1, yb1, zb1 := b1[0], b1[1], b1[2]
// 	xDiff0 := x0 == xb0
// 	yDiff0 := y0 == yb0
// 	zDiff0 := z0 == zb0
// 	xDiff1 := x1 == xb1
// 	yDiff1 := y1 == yb1
// 	zDiff1 := z1 == zb1
// 	if xDiff1 && yDiff1 && zDiff1 && xDiff0 && yDiff0 && zDiff0 {
// 		return true
// 	}
// 	return false
// }

func rotate(coords []int, mat [][]int) []int {
	x, y, z := coords[0], coords[1], coords[2]
	newX := mat[0][0]*x + mat[0][1]*x + mat[0][2]*x
	newY := mat[1][0]*y + mat[1][1]*y + mat[1][2]*y
	newZ := mat[2][0]*z + mat[2][1]*z + mat[2][2]*z
	return []int{newX, newY, newZ}
}
