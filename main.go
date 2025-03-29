package main

import (
	"fmt"
	"myFirstProject/highOrder"
)

func main() {

	fmt.Println("\n\n---- CalculateInvestments ----")
	fmt.Println("\n\nThis program calculates how much money you will have in the future based on a monthly investment and a fixed interest rate.")
	fmt.Println("It also calculates how long it will take to reach a target amount.")

	var principal, target, rate float64

	fmt.Print("\n\nEnter the initial amount: R$")
	fmt.Scanln(&principal)

	fmt.Print("Enter the target amount: R$")
	fmt.Scanln(&target)

	fmt.Print("Enter the rate (as a decimal):")
	fmt.Scanln(&rate)

	highOrder.CalculateInvestments(principal, target, rate)

}
