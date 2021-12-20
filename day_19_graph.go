package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/tabwriter"
)

type scanner struct {
	ID      int
	Beacons []*beacon
	Graph   *beaconGraph
	Coords  *scannerCoord
}

type scannerCoord struct {
	X, Y, Z int
}

func (s *scanner) GraphBeacons() {
	bg := newBeaconGraph(len(s.Beacons))
	for u := 0; u < len(s.Beacons)-1; u++ {
		b1 := s.Beacons[u]
		for v := u + 1; v < len(s.Beacons); v++ {
			b2 := s.Beacons[v]
			dist := b1.Distance(b2)
			bg.AddEdge(u, v, dist)
		}
	}
	s.Graph = bg
}

func (s *scanner) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "---[ Scanner %d ]---\n", s.ID)
	w := tabwriter.NewWriter(buf, 0, 0, 2, ' ', 0)
	for i, b1 := range s.Graph.Adj {
		fmt.Fprintf(w, "%d:\t", i)
		for _, dist := range b1 {
			fmt.Fprintf(w, "%.1f\t", dist)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buf.String()
}

type beacon struct {
	X, Y, Z int
}

func (b1 *beacon) Distance(b2 *beacon) float64 {
	x1, y1, z1 := float64(b1.X), float64(b1.Y), float64(b1.Z)
	x2, y2, z2 := float64(b2.X), float64(b2.Y), float64(b2.Z)

	x0 := math.Pow((x2 - x1), 2)
	y0 := math.Pow((y2 - y1), 2)
	z0 := math.Pow((z2 - z1), 2)

	return math.Sqrt(x0 + y0 + z0)
}

type beaconGraph struct {
	Nodes int
	Adj   [][]float64
}

func newBeaconGraph(beaconCount int) *beaconGraph {
	adjMatrix := make([][]float64, beaconCount)
	for i := 0; i < len(adjMatrix); i++ {
		adjMatrix[i] = make([]float64, beaconCount)
	}
	bg := &beaconGraph{
		Nodes: beaconCount,
		Adj:   adjMatrix,
	}
	return bg
}

func (bg *beaconGraph) AddEdge(u, v int, dist float64) {
	bg.Adj[u][v] = dist
	bg.Adj[v][u] = dist
}

func parseBeacons(f string) []*scanner {
	chunks := strings.Split(f, "\n\n")
	scanners := make([]*scanner, 0, len(chunks))

	for i, chunk := range chunks {
		beaconCoords := strings.Split(chunk, "\n")[1:]
		b := &scanner{
			ID:      i,
			Beacons: make([]*beacon, 0, len(beaconCoords)),
		}
		if i == 0 {
			b.Coords = &scannerCoord{0, 0, 0}
		}
		for _, coordsStr := range beaconCoords {
			coords := strings.Split(coordsStr, ",")
			xCoord, _ := strconv.Atoi(coords[0])
			yCoord, _ := strconv.Atoi(coords[1])
			zCoord, _ := strconv.Atoi(coords[2])
			scnr := &beacon{
				X: xCoord,
				Y: yCoord,
				Z: zCoord,
			}
			b.Beacons = append(b.Beacons, scnr)
		}
		scanners = append(scanners, b)
	}
	for _, s := range scanners {
		s.GraphBeacons()
	}
	return scanners
}
