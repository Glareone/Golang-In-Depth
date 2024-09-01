package main

import wwf "essentials.control-structure/working-with-files"

// variables available in the whole file
// in latest method it was replaced with real balance I read from the file
var dummyBalance float64 = 10000

// still used in functions
var dummyUserDeposit float64 = 0

func main() {
	var userBalance = wwf.GetBalanceFromFile()

	// UserChoiceHandler()
	// UserChoiceHandlerInfiniteLoop()
	// UserChoiceHandlerConditionalLoop()
	// UserChoiceHandlerInfiniteLoopSwitch()
	UserChoiceHandlerInfiniteLoopWithWriteToFile(userBalance)
}
