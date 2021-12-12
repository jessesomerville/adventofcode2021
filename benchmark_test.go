package main

import (
	"testing"
)

func BenchmarkSonarSweep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sonarSweep()
	}
}

func BenchmarkSonarSweepSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sonarSweepSlidingWindow()
	}
}

func BenchmarkDive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dive()
	}
}

func BenchmarkDiveWithAim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diveWithAim()
	}
}

func BenchmarkBinaryDiagnostic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryDiagnostic()
	}
}

func BenchmarkBinaryDiagnosticLifeSupport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryDiagnosticLifeSupport()
	}
}

func BenchmarkGiantSquid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		giantSquid()
	}
}

func BenchmarkGiantSquidLastWinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		giantSquidLastWinner()
	}
}

func BenchmarkHydrothermalVenture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hydrothermalVenture()
	}
}

func BenchmarkHydrothermalVentureDiagonals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hydrothermalVentureDiagonals()
	}
}

func BenchmarkLanternFish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lanternFish()
	}
}

func BenchmarkWhaleVsCrabs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		whaleVsCrabs()
	}
}

func BenchmarkWhaleVsCrabsGas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		whaleVsCrabsGas()
	}
}

func BenchmarkSevenSegment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sevenSegment()
	}
}

func BenchmarkSevenSegmentDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sevenSegmentDecode()
	}
}

func BenchmarkDecodeExperiment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeExperiment()
	}
}

func BenchmarkSmokeBasin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smokeBasin()
	}
}

func BenchmarkSmokeBasinLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smokeBasinLargest()
	}
}

func BenchmarkSyntaxScoring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syntaxScoring()
	}
}

func BenchmarkSyntaxScoringIncomplete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syntaxScoringIncomplete()
	}
}

func BenchmarkDumboOcto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dumboOctopus()
	}
}

func BenchmarkDumboOctoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dumboOctopusSync()
	}
}

func BenchmarkPassagePathing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		passagePathing()
	}
}

func BenchmarkPassagePathingRevisit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		passagePathingRevisit()
	}
}
