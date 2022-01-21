package idgenerator

import (
	"testing"
)

func TestGenerateDuplicateID(t *testing.T) {
	var x, y int64
	for i := 0; i < 10000; i++ {
		y = ID()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}
