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

// Array needs know the lenght in the declaration
func ReturnArray() [7]int {
	return [7]int{1, 2, 3, 4, 5, 6, 7}
}
