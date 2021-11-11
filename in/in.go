package in

func StringIn(array []string, needle string) (int, bool) {
	for i, item := range array {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}
