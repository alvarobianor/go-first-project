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
		seiLa          int    = 5
		seiLa2         string = "ALVIM REI"
		paralelepiledo []int  = []int{1, 2, 3, 4, 5}
	)
	fmt.Println("My favorite number is", u.GetRandomNumberAlv())
	fmt.Println("a test -> ", highOrder.Somar(idade)(44))
	fmt.Println("Print All things", nome)
	u.PrintAnything(seiLa, seiLa2, nome, idade, paralelepiledo, CONSTANTE)
}
