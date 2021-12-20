package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
)

var (
	//go:embed inputs/day_19.txt
	beaconsFile string

	//go:embed test_inputs/day_19.txt
	beaconsFileTest string
)

func beaconScanner() int {
	scanners := parseBeacons(beaconsFileTest)

	s0 := scanners[4]
	s1 := scanners[1]
	s1.Coords = &scannerCoord{
		X: 68,
		Y: -1246,
		Z: -43,
	}

	ol := sharedBeacons(s1, s0)
	coordsRelativeToRoot(s1, s0, ol)
	fmt.Println(s0.Coords)

	// s2 := scanners[2]
	// s3 := scanners[3]
	// s4 := scanners[4]
	// s1.Coords = &scannerCoord{
	// 	X: 68,
	// 	Y: -1246,
	// 	Z: -43,
	// }
	// s2.Coords = &scannerCoord{
	// 	X: 1105,
	// 	Y: -1205,
	// 	Z: 1229,
	// }
	// s3.Coords = &scannerCoord{
	// 	X: -92,
	// 	Y: -2380,
	// 	Z: -20,
	// }
	// s4.Coords = &scannerCoord{
	// 	X: -20,
	// 	Y: -1133,
	// 	Z: -1061,
	// }

	// findCoords(scanners, 0, map[int]bool{})

	// for i := 0; i < len(scanners)-1; i++ {
	// 	for j := i + 1; j < len(scanners); j++ {
	// 		a := scanners[i]
	// 		b := scanners[j]
	// 		if matches := sharedBeacons(a, b); matches != nil {
	// 			coordsRelativeToRoot(a, b, matches)
	// 		}
	// 	}
	// }

	// for _, s := range scanners {
	// 	fmt.Println(s.Coords)
	// }

	return 0
}

func findCoords(scanners []*scanner, idx int, found map[int]bool) {
	s0 := scanners[idx]

	for i := 0; i < len(scanners); i++ {
		if i == idx || found[i] {
			continue
		}
		s1 := scanners[i]
		if matches := sharedBeacons(s0, s1); matches != nil {
			coordsRelativeToRoot(s0, s1, matches)
			found[i] = true
			findCoords(scanners, i, found)
		}
	}
}

// If s0 and s1 have 12 overlapping beacons, return those beacons.  The return value is a slice
// whose subslices are []int{b0_ID, b1_ID} where s0_ID and s1_ID are the matching beacons.
// If they don't overlap, nil is returned.
func sharedBeacons(s0, s1 *scanner) [][]int {
	for _, s0BeaconDists := range s0.Graph.Adj {
		for _, s1BeaconDists := range s1.Graph.Adj {
			matching := make([][]int, 0, len(s0BeaconDists))
			for i, s0Dist := range s0BeaconDists {
				for j, s1Dist := range s1BeaconDists {
					if s0Dist == s1Dist {
						matching = append(matching, []int{i, j})
					}
				}
			}

			if len(matching) >= 12 {
				return matching
			}
		}
	}
	return nil
}

func (s *scanner) Distance(b *beacon) float64 {
	x1, y1, z1 := float64(s.Coords.X), float64(s.Coords.Y), float64(s.Coords.Z)
	x2, y2, z2 := float64(b.X), float64(b.Y), float64(b.Z)

	x0 := math.Pow((x2 - x1), 2)
	y0 := math.Pow((y2 - y1), 2)
	z0 := math.Pow((z2 - z1), 2)

	return math.Sqrt(x0 + y0 + z0)
}

func (s *scanner) DistanceS(b *scanner) float64 {
	x1, y1, z1 := float64(s.Coords.X), float64(s.Coords.Y), float64(s.Coords.Z)
	x2, y2, z2 := float64(b.Coords.X), float64(b.Coords.Y), float64(b.Coords.Z)

	x0 := math.Pow((x2 - x1), 2)
	y0 := math.Pow((y2 - y1), 2)
	z0 := math.Pow((z2 - z1), 2)

	return math.Sqrt(x0 + y0 + z0)
}

func coordsRelativeToRoot(s0, s1 *scanner, overlapping [][]int) {
	if s1.Coords != nil {
		return
		// log.Fatalf("Scanner ID=%d already has coordinates (%d, %d, %d)", s1.ID, s1.Coords.X, s1.Coords.Y, s1.Coords.Z)
	}
	if s0.Coords == nil {
		log.Fatalf("Scanner ID=%d does not have coordinates so Scanner ID=%d cannot be located", s0.ID, s1.ID)
	}

	var s0X, s0Y, s0Z, s1X, s1Y, s1Z []int
	for _, b := range overlapping {
		s0B := s0.Beacons[b[0]]
		s1B := s1.Beacons[b[1]]
		s0X = append(s0X, s0B.X)
		s0Y = append(s0Y, s0B.Y)
		s0Z = append(s0Z, s0B.Z)
		s1X = append(s1X, s1B.X)
		s1Y = append(s1Y, s1B.Y)
		s1Z = append(s1Z, s1B.Z)
	}

	var x, y, z int
	var xset, yset, zset bool
	s1Coords := [][]int{s1X, s1Y, s1Z}

	for _, coord := range s1Coords {
		if !xset {
			if newX, ok := findAllSameAdd(s0X, coord); ok {
				// fmt.Println("x plus", newX, i)
				x = newX
				// x = s0.Coords.X + newX
				xset = true
			} else if newX, ok := findAllSameSub(s0X, coord); ok {
				// fmt.Println("x minus", newX, i)
				x = newX
				// x = s0.Coords.X - newX
				xset = true
			}
		}
		if !yset {
			if newY, ok := findAllSameAdd(s0Y, coord); ok {
				// fmt.Println("y plus", newY, i)
				y = newY
				// y = s0.Coords.Y + newY
				yset = true
			} else if newY, ok := findAllSameSub(s0Y, coord); ok {
				// fmt.Println("y minus", newY, i)
				y = newY
				// y = s0.Coords.Y - newY
				yset = true
			}
		}
		if !zset {
			if newZ, ok := findAllSameAdd(s0Z, coord); ok {
				// fmt.Println("z plus", newZ, i)
				z = newZ
				// z = s0.Coords.Z + newZ
				zset = true
			} else if newZ, ok := findAllSameSub(s0Z, coord); ok {
				// fmt.Println("z minus", newZ, i)
				z = newZ
				// z = s0.Coords.Z - newZ
				zset = true
			}
		}
		if xset && yset && zset {
			break
		}
	}

	// xtmp := x + s0.Coords.X
	// ytmp := y + s0.Coords.Y
	// ztmp := z + s0.Coords.Z
	// c1 := s1.Beacons[overlapping[0][1]]
	// c2 := s0.Beacons[overlapping[0][0]]
	// c1.Y += x
	// c1.Z -= y
	// c1.X -= z

	// fmt.Println(c2)
	// fmt.Println(c1)

	fmt.Println()

	// fmt.Println(s0.ID, s1.ID)
	// fmt.Println(x, y, z)
	s1.Coords = &scannerCoord{
		X: s0.Coords.X - x,
		Y: s0.Coords.Y + y,
		Z: s0.Coords.Z - z,
	}
}

func findAllSameAdd(a, b []int) (int, bool) {
	baseline := a[0] + b[0]

	for i := 1; i < len(a); i++ {
		if a[i]+b[i] != baseline {
			return 0, false
		}
	}
	return baseline, true
}

func findAllSameSub(a, b []int) (int, bool) {
	baseline := a[0] - b[0]

	for i := 1; i < len(a); i++ {
		if a[i]-b[i] != baseline {
			return 0, false
		}
	}
	return baseline, true
}
