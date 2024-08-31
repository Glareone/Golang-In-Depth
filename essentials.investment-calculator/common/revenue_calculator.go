package common

import "math"

// Function name starts with uppercase to make it exportable
// Function name should start from Capital in order to be exported
func CalculateInvestmentAmountMultipleValues(investmentAmount int, years int, expectedRate float64, inflationRate float64) (float64, float64) {
	futureValue := float64(investmentAmount) * math.Pow(1+expectedRate/100, float64(years))
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	return futureValue, realFutureValue
}

// Declare and initialize futureValue and realFutureValues in function return declaration
// and in this case we dont need := which initializes the value, we can use just equal = sign
func CalculateInvestmentAmountMultipleValues2(investmentAmount int, years int, expectedRate float64, inflationRate float64) (futureValue float64, realFutureValue float64) {
	futureValue = float64(investmentAmount) * math.Pow(1+expectedRate/100, float64(years))
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, float64(years))

	// return futureValue, realFutureValue
	// we can use just return because we declared what we gonna return already
	// but better to explicitly say what we return
	return
}
