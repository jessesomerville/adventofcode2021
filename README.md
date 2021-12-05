# Advent of Code 2021

## Benchmarks

Benchmarks were performed in Windows Subsystem for Linux (WSL) on an i7-9700K CPU @ 3.60GHz.

### Day 01

```
BenchmarkSonarSweep-8                              29625             41149 ns/op           32768 B/op          1 allocs/op
BenchmarkSonarSweepSlidingWindow-8                 29902             40309 ns/op           32768 B/op          1 allocs/op
```

### Day 02

```
BenchmarkDive-8                                     3930            306415 ns/op           72496 B/op       2003 allocs/op
BenchmarkDiveWithAim-8                              3951            307060 ns/op           72496 B/op       2003 allocs/op
```

### Day 03

```
BenchmarkBinaryDiagnostic-8                         2340            508271 ns/op           56576 B/op       2006 allocs/op
BenchmarkBinaryDiagnosticLifeSupport-8              1941            613070 ns/op          271888 B/op      10644 allocs/op
```

### Day 04

```
BenchmarkGiantSquid-8                               4400            270575 ns/op          184953 B/op       2962 allocs/op
BenchmarkGiantSquidLastWinner-8                     3132            378273 ns/op          184948 B/op       2962 allocs/op
```