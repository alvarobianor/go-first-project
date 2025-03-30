package utils

import "fmt"

func ViewSlices() {
	slice := []int{}

	fmt.Println("Slice -> ", slice)
	fmt.Println("Slice Length -> ", len(slice))
	fmt.Println("Slice Capacity -> ", cap(slice))
}

func AddValueInSlice(slice *[]int, newValue int) {
	*slice = append(*slice, newValue)

}

func CreateSlice() ([]int, int) {
	sli := make([]int, 5, 50)
	sli[3] = 10
	fmt.Println("Slice -> ", sli)

	return sli, cap(sli)
}
