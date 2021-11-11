package pop

import "testing"

var stringQueue = []string{"one", "two", "three"}
var intQueue = []int{1, 2, 3}

func TestStringPop(t *testing.T) {
	poppedQueue := StringPop(stringQueue, -1)
	if len(poppedQueue) != 2 {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %d != '2'", len(poppedQueue))
	}
	if poppedQueue[0] != "one" {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %s != 'one'", poppedQueue[0])
	}
	if poppedQueue[1] != "two" {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %s != 'two'", poppedQueue[1])
	}
}

func TestStringPopByIndex(t *testing.T) {
	poppedQueue := StringPop(stringQueue, 1)
	if len(poppedQueue) != 2 {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %d != '2'", len(poppedQueue))
	}
	if poppedQueue[0] != "one" {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %s != 'one'", poppedQueue[0])
	}
	if poppedQueue[1] != "three" {
		t.Errorf("[func StringPop(array []string, index int) []string] -> %s != 'three'", poppedQueue[1])
	}
}

func BenchmarkStringPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPop(stringQueue, 0)
	}
}

func TestIntPop(t *testing.T) {
	poppedQueue := IntPop(intQueue, -1)
	if len(poppedQueue) != 2 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != '2'", len(poppedQueue))
	}
	if poppedQueue[0] != 1 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != 1", poppedQueue[0])
	}
	if poppedQueue[1] != 2 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != 2", poppedQueue[1])
	}
}

func TestIntPopByIndex(t *testing.T) {
	poppedQueue := IntPop(intQueue, 1)
	if len(poppedQueue) != 2 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != '2'", len(poppedQueue))
	}
	if poppedQueue[0] != 1 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != 1", poppedQueue[0])
	}
	if poppedQueue[1] != 3 {
		t.Errorf("[func IntPop(array []int, index int) []int] -> %d != 3", poppedQueue[1])
	}
}

func BenchmarkIntPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntPop(intQueue, 0)
	}
}
