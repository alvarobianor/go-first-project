package utils

import "fmt"

func CreateForOld(leng uint8) {
	fmt.Println("-----------------------")
	for i := 0; i < int(leng); i++ {
		fmt.Println("-> ", i)

	}
	fmt.Println("-----------------------")

}
