package int53IDGenerator

import (
	"testing"
)

func BenchmarkInt53IDGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt53IDGenerator().ID()
	}
}
