package utils

import "fmt"

func CreateForOld(leng uint8) {
	fmt.Println("-----------------------")
	for i := 0; i < int(leng); i++ {
		fmt.Println("-> ", i)

	}
	fmt.Println("-----------------------")

}

func AlternativeFo1(lenght uint) uint {
	var result uint
	i := 0

	for i < int(lenght) {
		result += uint(i)
		i++
	}
	return result
}
