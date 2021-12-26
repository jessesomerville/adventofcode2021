package week3

import (
	"math"
)

func (alignWith *scanner) align(toAlign *scanner, alignWithb, toAlignb int, intersect [][]int) {
	alignWithBeacon := alignWith.Beacons[alignWithb]
	toAlignBeacon := toAlign.Beacons[toAlignb]
	for _, beaconIdx := range intersect {
		if beaconIdx[0] == 0 {
			continue
		}

		s0Rel := alignWith.Beacons[beaconIdx[0]]
		dx0 := alignWithBeacon.X - s0Rel.X
		dy0 := alignWithBeacon.Y - s0Rel.Y
		dz0 := alignWithBeacon.Z - s0Rel.Z

		s1Rel := toAlign.Beacons[beaconIdx[1]]
		dx1 := toAlignBeacon.X - s1Rel.X
		dy1 := toAlignBeacon.Y - s1Rel.Y
		dz1 := toAlignBeacon.Z - s1Rel.Z

		if math.Abs(float64(dx0)) == math.Abs(float64(dy0)) || math.Abs(float64(dz0)) == math.Abs(float64(dy0)) || math.Abs(float64(dx0)) == math.Abs(float64(dz0)) {
			continue
		}

		rotationMat := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}

		if dx0 == dx1 {
			rotationMat[0][0] = 1
		} else if dx0 == -dx1 {
			rotationMat[0][0] = -1
		} else if dx0 == dy1 {
			rotationMat[1][0] = 1
		} else if dx0 == -dy1 {
			rotationMat[1][0] = -1
		} else if dx0 == dz1 {
			rotationMat[2][0] = 1
		} else if dx0 == -dz1 {
			rotationMat[2][0] = -1
		}

		if dy0 == dx1 {
			rotationMat[0][1] = 1
		} else if dy0 == -dx1 {
			rotationMat[0][1] = -1
		} else if dy0 == dy1 {
			rotationMat[1][1] = 1
		} else if dy0 == -dy1 {
			rotationMat[1][1] = -1
		} else if dy0 == dz1 {
			rotationMat[2][1] = 1
		} else if dy0 == -dz1 {
			rotationMat[2][1] = -1
		}

		if dz0 == dx1 {
			rotationMat[0][2] = 1
		} else if dz0 == -dx1 {
			rotationMat[0][2] = -1
		} else if dz0 == dy1 {
			rotationMat[1][2] = 1
		} else if dz0 == -dy1 {
			rotationMat[1][2] = -1
		} else if dz0 == dz1 {
			rotationMat[2][2] = 1
		} else if dz0 == -dz1 {
			rotationMat[2][2] = -1
		}

		for _, b := range toAlign.Beacons {
			b.rotate(rotationMat)
		}

		toAlign.Coords = &scannerCoord{
			X: alignWithBeacon.X - toAlignBeacon.X,
			Y: alignWithBeacon.Y - toAlignBeacon.Y,
			Z: alignWithBeacon.Z - toAlignBeacon.Z,
		}
		for _, b := range toAlign.Beacons {
			b.X += toAlign.Coords.X
			b.Y += toAlign.Coords.Y
			b.Z += toAlign.Coords.Z
		}
		return
	}
}

func (b *beacon) rotate(mat [][]int) {
	x, y, z := b.X, b.Y, b.Z
	newX := mat[0][0]*x + mat[1][0]*y + mat[2][0]*z
	newY := mat[0][1]*x + mat[1][1]*y + mat[2][1]*z
	newZ := mat[0][2]*x + mat[1][2]*y + mat[2][2]*z
	b.X = newX
	b.Y = newY
	b.Z = newZ
}
