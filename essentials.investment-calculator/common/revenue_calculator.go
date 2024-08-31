package common

import "math"

// Function name starts with uppercase to make it exportable
// Function name should start from Capital in order to be exported
func CalculateInvestmentAmountMultipleValues(investmentAmount int, years int, expectedRate float64, inflationRate float64) (float64, float64) {
	futureValue := float64(investmentAmount) * math.Pow(1+expectedRate/100, float64(years))
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	return futureValue, realFutureValue
}
