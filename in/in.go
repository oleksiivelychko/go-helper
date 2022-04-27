package in

func In[In comparable](array []In, needle In) (int, bool) {
	for i, item := range array {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}
