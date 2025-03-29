package highOrder

import (
	"fmt"
	"math"
	"strings"
)

func CalculateHowMonthsToMoney(currentMoney float64, rate float64, target float64) float32 {
	months := 0
	for currentMoney < target {
		currentMoney += currentMoney * (rate / 100)
		months++
	}
	return float32(months)

}

func CalculateHowTheTaxForMoney(currentMoney float64, target float64, months int64) string {
	tax := math.Pow((target/currentMoney), 1/float64(months)) - 1

	fomatedTax := tax * 100

	formattedTax := fmt.Sprintf("%.2f%%", fomatedTax)
	formattedTax = strings.Replace(formattedTax, ".", ",", 1)

	return formattedTax
}

func CalculateInvestments(money float64, target float64, rate float64) {

	fmt.Println("\n\n---- CalculateInvestments ----")
	fmt.Println("\nLet's start with the money: ", money, " and the target: ", int64(target))
	fmt.Println("The rate is: ", rate, "%")
	initialMoney := money
	var monthlyMoney float64 = 0.0
	month := 0
	for money <= target {
		monthlyMoney = money * (rate / 100)
		money += monthlyMoney
		month++
		fmt.Printf("\nMonth %d: R$%.2f -> + R$%.2f", month, money, monthlyMoney)
	}

	formattedSuffixYear := ""

	switch {
	case month >= 12 && month < 24:
		formattedSuffixYear = "year"
	case month >= 24:
		formattedSuffixYear = "years"
	case month < 12 && month > 1:
		formattedSuffixYear = "months"
	default:
		formattedSuffixYear = "month"
	}

	monthInYears := int(math.Floor(float64(month) / 12))
	monthsRestingInTheYear := month % 12

	formattedSuffixMonth := ""

	switch {
	case monthsRestingInTheYear == 1:
		formattedSuffixMonth = "month"
	case monthsRestingInTheYear > 1:
		formattedSuffixMonth = "months"
	}

	time := ""

	if month == 1 {
		time = "1 " + formattedSuffixYear
	} else if monthInYears == 1 && monthsRestingInTheYear == 0 {
		time = fmt.Sprintf("%d ", monthInYears) + formattedSuffixYear
	} else if monthInYears == 1 && monthsRestingInTheYear != 0 {
		time = fmt.Sprintf("%d ", monthInYears) + formattedSuffixYear + " and " + fmt.Sprintf("%d ", monthsRestingInTheYear) + formattedSuffixMonth
	} else if monthInYears > 1 && monthsRestingInTheYear == 0 {
		time = fmt.Sprintf("%d ", monthInYears) + formattedSuffixYear
	} else {
		time = fmt.Sprintf("%d ", monthInYears) + formattedSuffixYear + " and " + fmt.Sprintf("%d ", monthsRestingInTheYear) + formattedSuffixMonth
	}

	formattedRate := fmt.Sprintf("%.2f%%", rate)

	fmt.Println("\n\n---- Resume ----")
	fmt.Println("\nIn", time, "we can grow the money from R$", initialMoney, "to R$", int64(money), "with a rate of", strings.Replace(formattedRate, ".", ",", 1), "per month")

}
