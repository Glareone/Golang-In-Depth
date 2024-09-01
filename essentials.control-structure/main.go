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

	// UserChoiceHandler()
	UserChoiceHandlerLoop()
}
