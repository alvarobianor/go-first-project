package utils

import "fmt"

func ShowPointers() {
	fmt.Println("Pointers -> ")

	casa := "Rua Cel João Paracampos, 1075"

	var casaPointer *string = &casa

	fmt.Println("Casa -> ", casa)
	fmt.Println("CasaPointer -> ", casaPointer)
	fmt.Println("CasaPointer Value -> ", *casaPointer)
	fmt.Println("CasaPointer Address -> ", &casaPointer)
	fmt.Println("CasaPointer value will receive a new value")
	*casaPointer = "Rua Cel João Paracampos, 1075 - Casa 2"
	fmt.Println("CasaPointer Value -> ", *casaPointer)
	fmt.Println("Casa -> ", casa)

}

func Swap(a *int, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

func NewSap(a *int, b *int) (int, int) {
	temp := a
	a = b
	b = temp
	return *a, *b
}

func ExplainPointers(a *int) {
	fmt.Println("a is *int, lets see your values -> a, &a, *a", a, &a, *a)
	temp := &a
	fmt.Println("\n\nTemp is receiving the &a")
	fmt.Println("Temp -> ", temp)
	fmt.Println("Temp& -> ", &temp)
	fmt.Println("Temp* -> ", *temp)
	fmt.Println("\n\nTemp2 will receive the a")
	temp2 := a
	fmt.Println("Temp2-> ", temp2)
	fmt.Println("Temp2& -> ", &temp2)
	fmt.Println("Temp2* -> ", *temp2)
	fmt.Println("\n\nTemp3 will receive the *a")
	temp3 := *a
	fmt.Println("Temp3 -> ", temp3)
	fmt.Println("Temp3& -> ", &temp3)
	fmt.Println("Temp3* is invalid-> ")

	// a is a memoryadress of some variable, &a is the memoryadress of a, *a is the value of a
	// a, &a, *a (a == 0xc0000100e0) (&a == 0xc000068040) (*a == 10)
}

func CreatePointer(pointer *int, value int) *int {
	*pointer = value
	return pointer
}
