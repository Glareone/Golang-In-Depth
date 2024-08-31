package main

import (
	"fmt"
	"golang-essentials.com/investment-calculator/common"
	"math"
)

func main() {
	// initial_func()
	// investment_calculator_with_inflation()
	// explicit_type_annotation_calculator()
	// alternative_variable_declaration_calculator()
	// investment_calculator_with_inflation_user_input()
	investment_calculator_using_exported_function()
}

func initial_func() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}

func explicit_type_annotation_calculator() {
	var investmentAmount float64 = 1000 // explicitly declare the format
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}

func alternative_variable_declaration_calculator() {
	var investmentAmount, years float64 = 1000, 10 // explicitly declare the format
	// investmentAmount, years := 1000.0, 10.0 // implicit format declaration and without using var

	var string1, int1 = "10", 10 // another kind of declaration when you have different value types
	expectedReturnRate := 5.5    // := is a shortcut of var

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years) // same here, we dont use var, instead we use assignment using ":="
	fmt.Print(futureValue)
	fmt.Println("  string test  ", string1, int1)
}

func investment_calculator_with_inflation() {
	const investmentAmount float64 = 1000 // explicitly declare the format
	var expectedReturnRate = 5.5
	years := 10
	const inflationRate = 2.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureRealValueIncludingInflation = futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println("return rate with inflation: ", futureRealValueIncludingInflation)
}

func investment_calculator_with_inflation_user_input() {
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years int32
	var inflationRate float64 = 2.5

	fmt.Print("Investment Amount is: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate is: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureRealValueIncludingInflation = futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println("return rate without inflation: ", futureValue)
	fmt.Println("return rate with inflation: ", futureRealValueIncludingInflation)
}

func investment_calculator_using_exported_function() {
	var investmentAmount int = 10000
	var expectedReturnRate = 5.5
	var years int = 10
	var inflationRate float64 = 2.5

	var futureValue, realFutureValue = common.CalculateInvestmentAmountMultipleValues(investmentAmount, years, expectedReturnRate, inflationRate)
	fmt.Println("return rate without inflation: ", futureValue)
	fmt.Println("return rate with inflation: ", realFutureValue)
}
