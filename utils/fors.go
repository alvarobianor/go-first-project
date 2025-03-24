package utils

import "fmt"

func CreateForOld(leng uint8) {
	fmt.Println("-----------------------")
	for i := 0; i < int(leng); i++ {
		fmt.Println("-> ", i)

	}
	fmt.Println("-----------------------")

}

func AlternativeFor1(lenght uint8) uint8 {
	var result uint8 = 0
	i := 0

	for i < int(lenght) {
		result += uint8(i)
		i++
	}
	return result
}

func ForWithRange(elements []uint8) uint8 {
	var result uint8 = 0

	for _, elem := range elements {
		result += uint8(elem)
	}
	return result
}
