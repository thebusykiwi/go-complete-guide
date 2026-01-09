package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")
	
	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate) 
	
	fmt.Printf("Earings Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func outputText(field string) {
	fmt.Print(field)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	outputText(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio
}
