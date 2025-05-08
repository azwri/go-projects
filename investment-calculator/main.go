package main

import (
	"fmt"

	"github.com/azwri/go-projects/investment-calculator/utils"
)

func main() {
	investmentAmount := utils.PropmtUser("Investment amount")
	expectedReturnRate := utils.PropmtUser("Excepted Return Rate")
	years := utils.PropmtUser("Years")

	// fv (future value). rfv (real futrue value)
	fv, rfv := utils.CalculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println(fv)
	fmt.Println(rfv)
}
