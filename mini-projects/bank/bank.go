package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		checkBalance(accountBalance)
	} 
}


func checkBalance(accountBalance float64) {
	fmt.Println("Your current account balance: ", accountBalance)
}

func getUserInput(text string) float64 {
	var userInput float64
	printToConsole(text)
	fmt.Scan(&userInput)
	return userInput
}

func printToConsole(text string) {
	fmt.Print(text)
}
