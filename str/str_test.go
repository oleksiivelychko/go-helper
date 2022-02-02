package str

import "testing"

var intArray = []int{1, 2, 3, 4, 5}

func TestStringFromArrayInt(t *testing.T) {
	str := StringFromArrayInt(intArray, ",")
	if str != "1,2,3,4,5" {
		t.Errorf("[func StringFromArrayInt(slice []int, separator string) string] -> %s != '1,2,3,4,5'", str)
	}
}

func BenchmarkStringPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFromArrayInt(intArray, ",")
	}
}
