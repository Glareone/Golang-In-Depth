package main

import "fmt"

func UserChoiceHandler() {
	var userChoice int

	// if you input something unacceptable for this type - assignment will be ignored
	// and userChoice will still be 0 (default value for int)
	fmt.Scan(&userChoice)

	if userChoice == 1 {
		fmt.Println("Your balance is, ", dummyBalance)
	} else if userChoice == 2 {
		fmt.Print("How much do you want to deposit?")
		fmt.Scan(&dummyUserDeposit)

		if dummyUserDeposit <= 0 {
			fmt.Print("Invalid amount. Must be greater than 0.")
			return
		}

		dummyBalance += dummyUserDeposit
		fmt.Print("Balance updated: Now your balance is: ", dummyBalance)
	} else if userChoice == 3 {
		fmt.Println("How much do you want to withdraw?")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 || withdrawAmount > dummyUserDeposit {
			fmt.Print("Invalid amount. Must be greater than 0 and less than deposit sum")
			return
		}

		dummyBalance -= withdrawAmount
		fmt.Println("New balance is ", dummyBalance)
	} else {
		fmt.Println("You chose exit! Goodbye!")
	}
}

// using Infinite Loop
func UserChoiceHandlerInfiniteLoop() {
	// infinite loop using for {}
	for {
		var userChoice int

		// if you input something unacceptable for this type - assignment will be ignored
		// and userChoice will still be 0 (default value for int)
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Println("Your balance is, ", dummyBalance)
		} else if userChoice == 2 {
			fmt.Print("How much do you want to deposit?")
			fmt.Scan(&dummyUserDeposit)

			if dummyUserDeposit <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0.")
				continue
			}

			dummyBalance += dummyUserDeposit
			fmt.Print("Balance updated: Now your balance is: ", dummyBalance)
		} else if userChoice == 3 {
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > dummyUserDeposit {
				fmt.Print("Invalid amount. Must be greater than 0 and less than deposit sum")
				continue
			}

			dummyBalance -= withdrawAmount
			fmt.Println("New balance is ", dummyBalance)
		} else {
			fmt.Println("You chose exit! Goodbye!")
			break
		}
	}

	fmt.Println("Interaction with the bank ends! Thanks for choosing our bank")
}

func UserChoiceHandlerConditionalLoop() {
	var error error

	for error == nil {
		var userChoice int

		// if you input something unacceptable for this type - assignment will be ignored
		// and userChoice will still be 0 (default value for int)
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Println("Your balance is, ", dummyBalance)
		} else if userChoice == 2 {
			fmt.Print("How much do you want to deposit?")
			fmt.Scan(&dummyUserDeposit)

			if dummyUserDeposit <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0.")
				error = fmt.Errorf("incorrect deposit amount")
				continue
			}

			dummyBalance += dummyUserDeposit
			fmt.Print("Balance updated: Now your balance is: ", dummyBalance)
		} else if userChoice == 3 {
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > dummyUserDeposit {
				fmt.Print("Invalid amount. Must be greater than 0 and less than deposit sum")
				error = fmt.Errorf("incorrect withdraw amount")
				continue
			}

			dummyBalance -= withdrawAmount
			fmt.Println("New balance is ", dummyBalance)
		} else {
			fmt.Println("You chose exit! Goodbye!")
			break
		}
	}

	fmt.Println("Interaction with the bank ends! Thanks for choosing our bank")
}

func UserChoiceHandlerInfiniteLoopSwitch() {
loop:
	for {
		var userChoice int

		// if you input something unacceptable for this type - assignment will be ignored
		// and userChoice will still be 0 (default value for int)
		fmt.Scan(&userChoice)

		// break keyword is not needed for each case in Golang
		switch userChoice {
		case 1:
			fmt.Println("Your balance is, ", dummyBalance)
		case 2:
			fmt.Print("How much do you want to deposit?")
			fmt.Scan(&dummyUserDeposit)

			if dummyUserDeposit <= 0 {
				fmt.Print("Invalid amount. Must be greater than 0.")
				continue
			}

			dummyBalance += dummyUserDeposit
			fmt.Print("Balance updated: Now your balance is: ", dummyBalance)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > dummyUserDeposit {
				fmt.Print("Invalid amount. Must be greater than 0 and less than deposit sum")
				continue
			}

			dummyBalance -= withdrawAmount
			fmt.Println("New balance is ", dummyBalance)
		default:
			fmt.Println("You chose exit! Goodbye!")

			// labeled loop, you can use break with the labeled name of the loop to get out of the loop and continue the function
			break loop
		}
	}

	fmt.Println("Interaction with the bank ends! Thanks for choosing our bank")
}
