package main

import "fmt"

func main() {
	var revenue int32
	var expenses int32
	var taxRatePercent float64

	fmt.Print("Input Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input Tax Rate Percentage: ")
	fmt.Scan(&taxRatePercent)

	var earningsBeforeTax float64 = float64(revenue - expenses)
	var earningsAfterTax float64 = earningsBeforeTax * (1 - taxRatePercent/100)
	var earningsRatio = earningsBeforeTax / earningsAfterTax

	// formatted and not formatted output
	//fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	//fmt.Println("Earnings After Tax: ", earningsAfterTax)
	//fmt.Println("Earnings Ratio: ", earningsRatio)
	fmt.Printf("Earnings Before Tax: %v\nEarnings After Tax: %v\nEarnings Ratio: %v\n", earningsBeforeTax, earningsAfterTax, earningsRatio)
}
