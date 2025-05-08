package main

import (
	"github.com/azwri/go-projects/profit-calculator/utils"
)

func main() {

	revenue := utils.PromptUser("Revenue")
	expenses := utils.PromptUser("Expenses")
	taxRate := utils.PromptUser("Tax Rate")

	utils.CalclulateValues(revenue, expenses, taxRate)

}
