package main

import (
	"errors"
	"fmt"
)

func main() {
	var revenue, revenueErr = getUserInput("Input Revenue: ")
	var expenses, expensesError = getUserInput("Input Expenses: ")
	var taxRatePercent, taxRatePercentError = getUserInput("Input Tax Rate Percentage: ")

	if revenueErr != nil || expensesError != nil || taxRatePercentError != nil {
		if revenueErr != nil {
			fmt.Println("Error! ", revenueErr)
		}
		if expensesError != nil {
			fmt.Println("Error! ", expensesError)
		}
		if taxRatePercentError != nil {
			fmt.Println("Error! ", taxRatePercentError)
		}

		panic("Unhandled error occurred during input!")
	}

	var earningsBeforeTax float64 = float64(revenue - expenses)
	var earningsAfterTax float64 = earningsBeforeTax * (1 - taxRatePercent/100)
	var earningsRatio = earningsBeforeTax / earningsAfterTax

	// formatted and not formatted output
	//fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	//fmt.Println("Earnings After Tax: ", earningsAfterTax)
	//fmt.Println("Earnings Ratio: ", earningsRatio)

	// output formatted string, but without additional changes, just Values (because of %v)
	fmt.Printf("Earnings Before Tax: %v\nEarnings After Tax: %v\nEarnings Ratio: %v\n", earningsBeforeTax, earningsAfterTax, earningsRatio)

	// use %f formatting, float32 representation
	fmt.Printf("Earnings Before Tax: %f\nEarnings After Tax: %f\nEarnings Ratio: %f\n", earningsBeforeTax, earningsAfterTax, earningsRatio)
	// controlling how many decimals after dot we want to see. 2 decimals and 1
	fmt.Printf("Earnings Before Tax: %.2f\nEarnings After Tax: %.2f\nEarnings Ratio: %.1f\n", earningsBeforeTax, earningsAfterTax, earningsRatio)

	// multiline string formatting
	fmt.Printf(`Earnings Before Tax: %.2f

			Earnings After Tax: %.2f

			Earnings Ratio: %.1f`, earningsBeforeTax, earningsAfterTax, earningsRatio)

	// using Sprintf to format the string
	var earningsBeforeTaxFormattedOutput = fmt.Sprintf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	var earningsAfterTaxFormattedOutput = fmt.Sprintf("Earnings After Tax: %.2f\n", earningsAfterTax)
	var ratioFormattedOutput = fmt.Sprintf("Earnings Ratio: %.1f\n", earningsRatio)
	fmt.Print(earningsBeforeTaxFormattedOutput, earningsAfterTaxFormattedOutput, ratioFormattedOutput)
}

func getUserInput(formattedText string) (float64, error) {
	fmt.Println("Input Revenue: ")
	var userInput float64
	fmt.Scan(&userInput)

	if userInput < 0 {
		return 0, errors.New("user input is below zero, it was replaced with 0")
	}

	return userInput, nil
}
