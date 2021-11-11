package pop

func StringPop(array []string, index int) []string {
	if index < 0 {
		// pop off last item
		return array[:len(array)-1]
	} else {
		// ...or by index
		return append(array[:index], array[index+1:]...)
	}
}

func IntPop(array []int, index int) []int {
	if index < 0 {
		// pop off last item
		return array[:len(array)-1]
		// ...or by index
	} else {
		return append(array[:index], array[index+1:]...)
	}
}
