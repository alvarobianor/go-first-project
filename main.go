package main

import (
	"fmt"
	"myFirstProject/highOrder"
	u "myFirstProject/utils"
)

func main() {

	var idade, nome = 26, "Álvaro Bianor"
	const CONSTANTE = "ALVIM REI"
	var (
		seiLa          uint8  = 255
		seiLa2         string = "ALVIM REI"
		paralelepiledo []int  = []int{1, 2, 3, 4, 5}
	)
	fmt.Println("My favorite number is", u.GetRandomNumberAlv())
	fmt.Println("a test -> ", highOrder.Somar(idade)(44))
	fmt.Println("Print All things", nome)
	u.PrintAnything(seiLa, seiLa2, nome, idade, paralelepiledo, CONSTANTE)

	fmt.Println("ConvertInString -> ȹ (", u.ConvertInString(10084), ")")
	const integer = 10
	fmt.Println("TakeInt8 -> ȹ (", u.TakeInt8(integer), ")")
	fmt.Println("TakeInt16 -> ȹ (", u.TakeInt16(integer), ")")
}
