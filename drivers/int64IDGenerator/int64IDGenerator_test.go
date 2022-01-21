package int64IDGenerator

import (
	"testing"
)

func BenchmarkInt64IDGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt64IDGenerator().ID()
	}
}
