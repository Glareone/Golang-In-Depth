package main

import (
	wwf "essentials.control-structure/working-with-files"
	"fmt"
)

// variables available in the whole file
// in latest method it was replaced with real balance I read from the file
var dummyBalance float64 = 10000

// still used in functions
var dummyUserDeposit float64 = 0

func main() {
	var userBalance, err = wwf.GetBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR READING FILE")
		fmt.Println(err)
		fmt.Println("Operation cannot be proceed")
		// in this case we return from the application, we don't want to proceed.
		// but usually Go follows "Graceful" approach with error handling
		// it's also possible to use panic() command
		return
	}

	// UserChoiceHandler()
	// UserChoiceHandlerInfiniteLoop()
	// UserChoiceHandlerConditionalLoop()
	// UserChoiceHandlerInfiniteLoopSwitch()
	UserChoiceHandlerInfiniteLoopWithWriteToFile(userBalance)
}
