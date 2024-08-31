package main

import "fmt"

// variables available in the whole file
var dummyBalance float64 = 10000
var dummyUserDeposit float64 = 0

func main() {
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1: Check the balance")
	fmt.Println("2: Deposit money")
	fmt.Println("3: Draw money")
	fmt.Println("4: Exit the application")

	fmt.Print("Your choice is: ")
	var userChoice int

	// if you input something unacceptable for this type - assignment will be ignored
	// and userChoice will still be 0 (default value for int)
	fmt.Scan(&userChoice)

	if userChoice == 1 {
		fmt.Println("Your balance is, ", dummyBalance)
	} else if userChoice == 2 {
		fmt.Print("How much do you want to deposit?")
		fmt.Scan(&dummyUserDeposit)
		dummyBalance += dummyUserDeposit
		fmt.Print("Balance updated: Now your balance is: ", dummyBalance)
	} else if userChoice == 3 {
		fmt.Println("How much do you want to withdraw?")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		dummyBalance -= withdrawAmount
		fmt.Println("New balance is ", dummyBalance)
	} else {
		fmt.Println("You chose exit! Goodbye!")
	}
}
