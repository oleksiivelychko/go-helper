package in

import "testing"

var array = []string{"one", "two", "three"}

func TestStringIn(t *testing.T) {
	index, result := StringIn(array, "one")
	if !result {
		t.Errorf("[func StringIn(array []string, val string) (int, bool)] -> %d != '0'", index)
	}
}

func BenchmarkStringIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringIn(array, "one")
	}
}
