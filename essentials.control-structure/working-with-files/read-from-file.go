package working_with_files

import (
	"errors"
	"os"
	"strconv"
)

const predefinedBalance float64 = 1000

// GetBalanceFromFile 
// only Uppercase named function will be exported and could be used in other places of the application
func GetBalanceFromFile() (float64, error) {
	// I declare error variable as underscore _ to explicitly say i dont want to use it
	userBalanceBytes, err := os.ReadFile(declaredFileName)

	if err != nil {
		// gracefully handle the occurred and continue
		return predefinedBalance, errors.New("file cannot be read")
	}

	balanceText := string(userBalanceBytes)

	// error could be omit using _ underscore
	var userBalance, parseBalanceError = strconv.ParseFloat(balanceText, 64)

	if parseBalanceError != nil {
		// we don't want to proceed with command, and we use command to stop application execution!
		// application crashes
		panic("balance from the file cannot be converted properly")
	}

	return userBalance, nil
}
