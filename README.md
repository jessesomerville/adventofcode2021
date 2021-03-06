# Advent of Code 2021

Solutions to all 25 days (except day 23 which I solved by hand) of
[Advent of Code](https://adventofcode.com) 2021 written in Go.

## Benchmarks

> **NOTE:** I stopped doing benchmarks after day 12 because it became more trouble than it was worth.

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
BenchmarkBinaryDiagnostic-8                         5340            227984 ns/op           16513 B/op          3 allocs/op
BenchmarkBinaryDiagnosticLifeSupport-8              3246            371632 ns/op          101617 B/op       3549 allocs/op
```

### Day 04

```
BenchmarkGiantSquid-8                               4400            270575 ns/op          184953 B/op       2962 allocs/op
BenchmarkGiantSquidLastWinner-8                     3132            378273 ns/op          184948 B/op       2962 allocs/op
```

### Day 05 (after manually moving inputs)

```
BenchmarkHydrothermalVenture-8                       825           1250346 ns/op         8003626 B/op          1 allocs/op
BenchmarkHydrothermalVentureDiagonals-8              615           1801328 ns/op         8003605 B/op          1 allocs/op
```

### Day 06

```
BenchmarkLanternFish-8                            104688             11389 ns/op           25344 B/op        257 allocs/op
```

### Day 07

```
BenchmarkWhaleVsCrabs-8                            15492             78772 ns/op           24600 B/op          3 allocs/op
BenchmarkWhaleVsCrabsGas-8                         56683             20940 ns/op           24576 B/op          2 allocs/op
```

### Day 08

```
BenchmarkSevenSegment-8                            10000            119339 ns/op           65792 B/op        802 allocs/op
BenchmarkSevenSegmentDecode-8                       5180            235132 ns/op           76992 B/op       1202 allocs/op
```

### Day 09

```
BenchmarkSmokeBasin-8                               8551            137188 ns/op               0 B/op          0 allocs/op
BenchmarkSmokeBasinLargest-8                        1014           1160128 ns/op          277790 B/op       1381 allocs/op
```

### Day 10

```
BenchmarkSyntaxScoring-8                            9054            139879 ns/op           15128 B/op         13 allocs/op
BenchmarkSyntaxScoringIncomplete-8                  7232            167333 ns/op           16256 B/op        459 allocs/op
```

### Day 11

```
BenchmarkDumboOcto-8                                8467            139225 ns/op          109760 B/op       1921 allocs/op
BenchmarkDumboOctoSync-8                            2178            526448 ns/op          418602 B/op       7133 allocs/op
```

### Day 12

```
BenchmarkPassagePathing-8                            826           1391540 ns/op            6513 B/op         95 allocs/op
BenchmarkPassagePathingRevisit-8                      21          53707319 ns/op            6499 B/op         95 allocs/op
```