package highOrder

import (
	"fmt"
	"strconv"
)

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

func Alvaro() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}
