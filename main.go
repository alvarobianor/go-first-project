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
		paralelepiledo [5]int = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("My favorite number is", u.GetRandomNumberAlv())
	fmt.Println("a test -> ", highOrder.Somar(idade)(44))
	fmt.Println("Print All things", nome)
	u.PrintAnything(seiLa, seiLa2, nome, idade, paralelepiledo, CONSTANTE)

	fmt.Println("ConvertInString -> ȹ (", u.ConvertInString(10084), ")")
	const integer = 10
	fmt.Println("TakeInt8 -> ȹ (", u.TakeInt8(integer), ")")
	fmt.Println("TakeInt16 -> ȹ (", u.TakeInt16(integer), ")")
	fmt.Println("ConstantArray -> ȹ (", u.ConstantArray(), ")")
	fmt.Println(u.PrintFuria())
	fmt.Println(u.DividirCSGO(131, 3))
	fmt.Println(highOrder.ConvertIntInString(1038))

	u.CreateForOld(4)

	fmt.Println("ForWithRange -> Its a Sum(", u.ForWithRange([]uint8{1, 2, 3, 4, 29}), ")")
	u.ForInPointers()
	u.Conditional(3)

	u.WhatTypeIsThis('k')
	numbersOfDefers := u.DoDefer()
	fmt.Println("Number of defers ->", numbersOfDefers)
	u.UndestandingDefers(80)

	a, b := 10, 20
	fmt.Println("Before Swap -> a, b", a, b)
	a, b = u.NewSap(&a, &b)

	fmt.Println("After Swap -> a, b", a, b)

	pointValue := 10
	var pointer *int = nil
	fmt.Println("Pointer is nil -> ", pointer)
	pointer = u.CreatePointer(&pointValue, 100)
	fmt.Println("Pointer is not nil, and his value -> ", pointer, ", ", *pointer)
	u.ViewSlices()
	sli := []int{}
	u.AddValueInSlice(&sli, 10)
	u.AddValueInSlice(&sli, 10)
	u.AddValueInSlice(&sli, 10)
	u.AddValueInSlice(&sli, 10)
	fmt.Println("Slice -> ", sli)

	newSlice := make([]int, 5, 10)
	fmt.Println("New Slice ->", newSlice, "Length ->", len(newSlice), "Capacity ->", cap(newSlice))

}
