package main

import (
	"fmt"
	"myFirstProject/highOrder"
	u "myFirstProject/utils"
)

func main() {
	fmt.Println("My favorite number is", u.GetRandomNumberAlv())
	fmt.Println("a test -> ", highOrder.Somar(2)(44))
	// fmt.Println("Print All things", u.PrintAnything(1, 3, 4,5, 6,64, 32,3 ,1, "as", '1', 20.4))
}

