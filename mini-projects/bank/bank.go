package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	
	for {
		fmt.Println("What do you want to do?")
	
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit amount")
		fmt.Println("3. Withdraw amount")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		if choice == 1 {
			checkBalance(accountBalance)
		} else if choice == 2 {
			accountBalance = depositAmount(accountBalance)
		} else if choice == 3 {
			accountBalance = withdrawAmount(accountBalance)
		} else {
			printToConsole("Thank you for choosing us!\n")
			return
		}
	}
}


func checkBalance(accountBalance float64) {
	fmt.Println("Your current account balance:", accountBalance)
}

func depositAmount(accountBalance float64) float64 {
	userDeposit := getUserInput("Enter your amount: ")
	if userDeposit <= 0 {
		printToConsole("Invalid amount. Must be greater than 0.\n")
		return accountBalance
	}
	accountBalance += userDeposit
	printToConsole("Amount added successfully\n")
	checkBalance(accountBalance)
	return accountBalance
}

func withdrawAmount(accountBalance float64) float64 {
	userWithdraw := getUserInput("Enter your amount: ") 

	if userWithdraw <= 0 {
		printToConsole("Invalid amount. Must be greater than 0.\n")
		return accountBalance
	} 
		
	if userWithdraw > accountBalance {
		printToConsole("Insufficient funds\n")
		return accountBalance
	}

	accountBalance -= userWithdraw
	printToConsole("Amount withdrawn successfully\n")
	checkBalance(accountBalance)

	return accountBalance
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
