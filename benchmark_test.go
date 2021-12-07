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
