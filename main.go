package main

import (
	"fmt"
	u "myFirstProject/utils"
)

func main() {
	fmt.Println("My favorite number is", u.GetRandomNumberAlv())
	fmt.Println("Print All things", u.PrintAnything(1, 3, 4,5, 6,64, 32,3 ,1, "as", '1', 20.4))
}

