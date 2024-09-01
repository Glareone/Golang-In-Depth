package working_with_files

import (
	"os"
	"strconv"
)

func GetBalanceFromFile() float64 {
	// I declare error variable as underscore _ to explicitly say i dont want to use it
	userBalanceBytes, _ := os.ReadFile(declaredFileName)
	balanceText := string(userBalanceBytes)

	// here again I dont want to receive errors from the method, instead I omit it using underscore _
	userBalance, _ := strconv.ParseFloat(balanceText, 64)

	return userBalance
}
