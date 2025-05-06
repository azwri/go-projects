package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// Earning before tax
	ebt := revenue - expenses

	// Profit
	profit := ebt * (1 - taxRate/100)

	// Ratio
	ratio := ebt / profit

	// output
	fmt.Printf("Earning before tax: %.2f.\nProfit: %.2f.\nRatio: %.2f.\n", ebt, profit, ratio)

}
