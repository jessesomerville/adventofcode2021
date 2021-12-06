package main

import "testing"

func TestSonarSweep(t *testing.T) {
	want := 1475
	if got := sonarSweep(); got != want {
		t.Errorf("sonarSweep() = %d, want = %d", got, want)
	}
}

func TestSonarSweepSlidingWindow(t *testing.T) {
	want := 1516
	if got := sonarSweepSlidingWindow(); got != want {
		t.Errorf("sonarSweepSlidingWindow() = %d, want = %d", got, want)
	}
}

func TestDive(t *testing.T) {
	want := 1990000
	if got := dive(); got != want {
		t.Errorf("dive() = %d, want = %d", got, want)
	}
}

func TestDiveWithAim(t *testing.T) {
	want := 1975421260
	if got := diveWithAim(); got != want {
		t.Errorf("diveWithAim() = %d, want = %d", got, want)
	}
}

func TestBinaryDiagnostic(t *testing.T) {
	want := 4191876
	if got := binaryDiagnostic(); got != want {
		t.Errorf("binaryDiagnostic() = %d, want = %d", got, want)
	}
}

func TestBinaryDiagnosticLifeSupport(t *testing.T) {
	want := 3414905
	if got := binaryDiagnosticLifeSupport(); got != want {
		t.Errorf("binaryDiagnosticLifeSupport() = %d, want = %d", got, want)
	}
}

func TestGiantSquid(t *testing.T) {
	want := 41503
	if got := giantSquid(); got != want {
		t.Errorf("giantSquid() = %d, want = %d", got, want)
	}
}

func TestGiantSquidLastWinner(t *testing.T) {
	want := 3178
	if got := giantSquidLastWinner(); got != want {
		t.Errorf("giantSquidLastWinner() = %d, want = %d", got, want)
	}
}

func TestLanternFish(t *testing.T) {
	want := 1675781200288
	if got := lanternFish(); got != want {
		t.Errorf("lanternFish() = %d, want = %d", got, want)
	}
}
