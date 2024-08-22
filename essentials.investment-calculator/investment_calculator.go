package main

import (
	"fmt"
	"math"
)

func main() {
	initial_func()
	explicit_type_annotation_calculator()
	alternative_variable_declaration_calculator()
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
