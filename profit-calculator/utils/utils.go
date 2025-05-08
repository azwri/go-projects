package utils

import (
	"fmt"
	"log"
)

func PromptUser(text string) float64 {
	fmt.Print(text + ": ")
	var value float64
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatal("enter error, we expect numbers only")
	}
	return value
}

func CalclulateValues(revenue, expenses, taxRate float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earning before tax: %.2f.\nProfit: %.2f.\nRatio: %.2f.\n", ebt, profit, ratio)
}
