package highOrder

import "strconv"

func Somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func ConvertIntInString(a int) string {
	return strconv.FormatInt(int64(a), 10)
}
