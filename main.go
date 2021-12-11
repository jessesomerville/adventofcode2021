package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func main() {
	// sonarSweep()
	// sonarSweepSlidingWindow()
	// dive()
	// diveWithAim()
	// binaryDiagnostic()
	// binaryDiagnosticLifeSupport()
	// giantSquid()
	// giantSquidLastWinner()
	// hydrothermalVenture()
	// hydrothermalVentureDiagonals()
	// lanternFish()
	// whaleVsCrabs()
	// whaleVsCrabsGas()
	// sevenSegment()
	// sevenSegmentDecode()
	// smokeBasin()
	// smokeBasinLargest()
	// syntaxScoring()
	// syntaxScoringIncomplete()
	// dumboOctopus()
	fmt.Println(dumboOctopusSync())
}
