package cmp

import "testing"

func TestStringCmp(t *testing.T) {
	isEqual := StringCmp("one", "one")
	if !isEqual {
		t.Errorf("[func StringCmp(item1 string, item2 string) bool] -> %t != 'false'", isEqual)
	}
}

func BenchmarkStringCmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringCmp("one", "one")
	}
}
