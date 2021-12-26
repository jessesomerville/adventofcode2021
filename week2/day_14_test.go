package week2

import (
	_ "embed"
	"testing"
)

//go:embed day_14.txt
var polymerFile string

func BenchmarkPolymerization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Polymerization(polymerFile)
	}
}

// BenchmarkPolymerization-8   	       2	 511693616 ns/op	159511108 B/op	    1649 allocs/op
// BenchmarkPolymerization-8   	       4	 313847284 ns/op	159531776 B/op	    3616 allocs/op
// BenchmarkPolymerization-8   	       5	 234047220 ns/op	159572715 B/op	    7454 allocs/op
