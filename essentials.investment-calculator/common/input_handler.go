package common

import "fmt"

func HandleInput() (investmentAmount int, years int, expectedReturnRate float64, inflationRate float64) {
	fmt.Print("Investment Amount is: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate is: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Inflation Rate is: ")
	fmt.Scan(&inflationRate)

	return investmentAmount, years, expectedReturnRate, inflationRate
}
