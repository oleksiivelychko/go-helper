package in

import "testing"

var array = []string{"one", "two", "three"}

func TestStringIn(t *testing.T) {
	index, result := In(array, "one")
	if !result {
		t.Errorf("[func In[In comparable](array []In, needle In) (int, bool)] -> %d != '0'", index)
	}
}

func BenchmarkStringIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		In(array, "one")
	}
}
