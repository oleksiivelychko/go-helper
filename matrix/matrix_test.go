package matrix

import (
	"math/rand"
	"testing"
)

func TestIntMatrix(t *testing.T) {
	matrix := IntMatrix(5, 5)
	if len(matrix) != 5 {
		t.Errorf("[func IntMatrix(rows, cols int) [][]int] -> %d != '5'", len(matrix))
	}
	if cap(matrix) != 5 {
		t.Errorf("[func IntMatrix(rows, cols int) [][]int] -> %d != '5'", cap(matrix))
	}
}

func BenchmarkIntMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMatrix(rand.Intn(100), rand.Intn(100))
	}
}
