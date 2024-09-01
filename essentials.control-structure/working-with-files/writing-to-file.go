package working_with_files

import (
	"fmt"
	"os"
)

func WriteBalanceToFile(balance float64) {
	balanceAsString := fmt.Sprint(balance)
	// 0644 - your current user can read and write, other users can read only
	os.WriteFile(declaredFileName, []byte(balanceAsString), 0644)
}
