package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.8

func main() {
	var investmentAmount, years float64
	var expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Printf("Future Value: %v\n", futureValue)
	// fmt.Printf("Future Real Value: %.2f\n", futureRealValue)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(field string) {
	fmt.Print(field)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv := fv / math.Pow(1 + inflationRate / 100, years)
	return fv, frv
}