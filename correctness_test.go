package main

import (
	"testing"

	"github.com/jessesomerville/adventofcode2021/week1"
	"github.com/jessesomerville/adventofcode2021/week2"
)

func TestSonarSweep(t *testing.T) {
	want := 1475
	if got := week1.SonarSweep(depthsFile); got != want {
		t.Errorf("sonarSweep() = %d, want = %d", got, want)
	}
}

func TestSonarSweepSlidingWindow(t *testing.T) {
	want := 1516
	if got := week1.SonarSweepSlidingWindow(depthsFile); got != want {
		t.Errorf("sonarSweepSlidingWindow() = %d, want = %d", got, want)
	}
}

func TestDive(t *testing.T) {
	want := 1990000
	if got := week1.Dive(commandsFile); got != want {
		t.Errorf("dive() = %d, want = %d", got, want)
	}
}

func TestDiveWithAim(t *testing.T) {
	want := 1975421260
	if got := week1.DiveWithAim(commandsFile); got != want {
		t.Errorf("diveWithAim() = %d, want = %d", got, want)
	}
}

func TestBinaryDiagnostic(t *testing.T) {
	want := 4191876
	if got := week1.BinaryDiagnostic(diagFile); got != want {
		t.Errorf("binaryDiagnostic() = %d, want = %d", got, want)
	}
}

func TestBinaryDiagnosticLifeSupport(t *testing.T) {
	want := 3414905
	if got := week1.BinaryDiagnosticLifeSupport(diagFile); got != want {
		t.Errorf("binaryDiagnosticLifeSupport() = %d, want = %d", got, want)
	}
}

func TestGiantSquid(t *testing.T) {
	want := 41503
	if got := week1.GiantSquid(bingoFile); got != want {
		t.Errorf("giantSquid() = %d, want = %d", got, want)
	}
}

func TestGiantSquidLastWinner(t *testing.T) {
	want := 3178
	if got := week1.GiantSquidLastWinner(bingoFile); got != want {
		t.Errorf("giantSquidLastWinner() = %d, want = %d", got, want)
	}
}

func TestHydrothermalVenture(t *testing.T) {
	want := 7468
	if got := week1.HydrothermalVenture(); got != want {
		t.Errorf("hydrothermalVenture() = %d, want = %d", got, want)
	}
}

func TestHydrothermalVentureDiagonals(t *testing.T) {
	want := 22364
	if got := week1.HydrothermalVentureDiagonals(); got != want {
		t.Errorf("hydrothermalVentureDiagonals() = %d, want = %d", got, want)
	}
}

func TestLanternFish(t *testing.T) {
	want := 1675781200288
	if got := week1.LanternFish(lanternFishFile); got != want {
		t.Errorf("lanternFish() = %d, want = %d", got, want)
	}
}

func TestWnaleVsCrabs(t *testing.T) {
	want := 347011
	if got := week1.WhaleVsCrabs(crabsFile); got != want {
		t.Errorf("whaleVsCrabs() = %d, want = %d", got, want)
	}
}

func TestWnaleVsCrabsGas(t *testing.T) {
	want := 98363777
	if got := week1.WhaleVsCrabsGas(crabsFile); got != want {
		t.Errorf("whaleVsCrabs() = %d, want = %d", got, want)
	}
}

func TestSevenSegment(t *testing.T) {
	want := 272
	if got := week2.SevenSegment(segmentsFile); got != want {
		t.Errorf("sevenSegment() = %d, want = %d", got, want)
	}
}

func TestSevenSegmentDecode(t *testing.T) {
	want := 1007675
	if got := week2.SevenSegmentDecode(segmentsFile); got != want {
		t.Errorf("sevenSegmentDecode() = %d, want = %d", got, want)
	}
}

func TestSmokeBasin(t *testing.T) {
	want := 570
	if got := week2.SmokeBasin(heightMapFile); got != want {
		t.Errorf("smokeBasin() = %d, want = %d", got, want)
	}
}

func TestSmokeBasinLargest(t *testing.T) {
	want := 899392
	if got := week2.SmokeBasinLargest(heightMapFile); got != want {
		t.Errorf("smokeBasinLargest() = %d, want = %d", got, want)
	}
}

func TestSyntaxScoring(t *testing.T) {
	want := 240123
	if got := week2.SyntaxScoring(chunksFile); got != want {
		t.Errorf("syntaxScoring() = %d, want = %d", got, want)
	}
}

func TestSyntaxScoringIncomplete(t *testing.T) {
	want := 3260812321
	if got := week2.SyntaxScoringIncomplete(chunksFile); got != want {
		t.Errorf("syntaxScoringIncomplete() = %d, want = %d", got, want)
	}
}

func TestDumboOcto(t *testing.T) {
	want := 1571
	if got := week2.DumboOctopus(octosFile); got != want {
		t.Errorf("dumboOctopus() = %d, want = %d", got, want)
	}
}

func TestDumboOctoSync(t *testing.T) {
	want := 387
	if got := week2.DumboOctopusSync(octosFile); got != want {
		t.Errorf("dumboOctopusSync() = %d, want = %d", got, want)
	}
}

func TestPassagePathing(t *testing.T) {
	want := 4691
	if got := week2.PassagePathing(passageFile); got != want {
		t.Errorf("passagePathing() = %d, want = %d", got, want)
	}
}

func TestPassagePathingRevisit(t *testing.T) {
	want := 140718
	if got := week2.PassagePathingRevisit(passageFile); got != want {
		t.Errorf("passagePathingRevisit() = %d, want = %d", got, want)
	}
}
