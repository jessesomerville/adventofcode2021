package main

import (
	_ "embed"
	"fmt"
	"math"
)

var (
	//go:embed inputs/day_19.txt
	beaconsFile string
)

func beaconScanner() int {
	scanners := parseBeacons(beaconsFile)

	findCoords(scanners, 0, map[int]bool{0: true})
	beaconMap := map[string]bool{}
	for _, s := range scanners {
		for _, b := range s.Beacons {
			bCoordStr := fmt.Sprint(b)
			if !beaconMap[bCoordStr] {
				beaconMap[bCoordStr] = true
			}
		}
	}

	return len(beaconMap)
}

func beaconScannerMaxDist() int {
	scanners := parseBeacons(beaconsFile)

	findCoords(scanners, 0, map[int]bool{0: true})

	maxDist := 0
	for i, s0 := range scanners {
		for j, s1 := range scanners {
			if i == j {
				continue
			}
			if dist := s0.ManhattanDist(s1); dist > maxDist {
				maxDist = dist
			}
		}
	}

	fmt.Println(maxDist)

	return 0
}

func findCoords(scanners []*scanner, idx int, found map[int]bool) {
	s0 := scanners[idx]

	for i := 0; i < len(scanners); i++ {
		if i == idx || found[i] {
			continue
		}
		s1 := scanners[i]
		if s0i, s1j, intersect := sharedBeacons(s0, s1); intersect != nil {
			s0.align(s1, s0i, s1j, intersect)
			found[i] = true
			findCoords(scanners, i, found)
		}
	}
}

func sharedBeacons(s0, s1 *scanner) (i, j int, intersect [][]int) {
	for i, s0BeaconDists := range s0.Graph.Adj {
		for j, s1BeaconDists := range s1.Graph.Adj {
			matching := make([][]int, 0, len(s0BeaconDists))
			for i, s0Dist := range s0BeaconDists {
				for j, s1Dist := range s1BeaconDists {
					if s0Dist == s1Dist {
						matching = append(matching, []int{i, j})
					}
				}
			}

			if len(matching) >= 11 {
				return i, j, matching
			}
		}
	}
	return 0, 0, nil
}

func (s0 *scanner) ManhattanDist(s1 *scanner) int {
	dx := math.Abs(float64(s0.Coords.X - s1.Coords.X))
	dy := math.Abs(float64(s0.Coords.Y - s1.Coords.Y))
	dz := math.Abs(float64(s0.Coords.Z - s1.Coords.Z))
	return int(dx + dy + dz)
}
