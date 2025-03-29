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
