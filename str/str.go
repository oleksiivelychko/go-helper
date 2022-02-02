package str

import "strconv"

func StringFromArrayInt(slice []int, separator string) string {
	str := ""
	for _, v := range slice {
		if len(str) > 0 {
			str += separator
		}
		str += strconv.Itoa(v)
	}
	return str
}
