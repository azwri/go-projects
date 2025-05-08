// Package utils provides utility functions for financial calculations.
//
// This package includes functions for calculating future values of investments
// and for prompting user input in a command-line interface.
package utils

import (
	"fmt"
	"log"
	"math"
)

const inflationRate = 2.5

// CalculateFutureValues calculates the future value (FV) and real future value (RFV) of an investment
// based on the initial investment amount, expected annual return rate, and the number of years.
//
// Parameters:
//   - investentAmount: The initial amount of money invested.
//   - exepectedReturnRate: The expected annual return rate as a percentage.
//   - years: The number of years the investment will grow.
//
// Returns:
//   - fv: The future value of the investment after the specified number of years.
//   - rfv: The real future value of the investment, adjusted for inflation.
//
// Note: The function assumes the presence of a global variable `inflationRate`
// representing the annual inflation rate as a percentage.
func CalculateFutureValues(investentAmount, exepectedReturnRate, years float64) (float64, float64) {
	fv := investentAmount * math.Pow(1+exepectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

// PropmtUser prompts the user with the provided text, reads a numeric input from the console,
// and returns it as a float64. If the input is invalid or not a number, the function logs a fatal error
// and terminates the program.
//
// Parameters:
//   - text: A string representing the prompt message to display to the user.
//
// Returns:
//   - float64: The numeric value entered by the user.
func PropmtUser(text string) float64 {
	var value float64
	fmt.Print(text + ": ")
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatal("error input, you have to enter numbers")
	}
	return value
}
