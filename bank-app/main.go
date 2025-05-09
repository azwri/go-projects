package main

import (
	"fmt"

	"github.com/azwri/go-projects/bank-app/utils"
)

func main() {
	fmt.Println("Welcome to Bank App")

	for {
		balance := utils.ReadBalanceFromFile()
		fmt.Println(`Choose service: 
----------------------------------------
----------------------------------------
1. Check balance.
2. Deposite money.
3. withdraw money.
4. Exit
----------------------------------------
----------------------------------------
		`)
		var choice int
		fmt.Print("Your choice: -> ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %v\n", balance)
		case 2:
			var money float64
			fmt.Print("Deposite money: -> ")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("Error, deposite money should be greater than ZERO")
			} else {
				balance += money
				utils.WriteBalanceToFile(balance)
				fmt.Printf("Your balance is %v\n", balance)
			}
		case 3:
			var money float64
			fmt.Print("Withdrawal money: -> ")
			fmt.Scan(&money)
			if money <= 0 || money > balance {
				fmt.Println("Withdrawal money is less than ZERO, or greater than your balance!")
				fmt.Printf("Your balance is %v\n", balance)
			} else {
				balance -= money
				utils.WriteBalanceToFile(balance)
				fmt.Printf("Your balance is %v\n", balance)
			}
		case 4:
			fmt.Println("Thank you for using Bank App!")
			return
		default:
			fmt.Println("Error choice! choose one of the below numbers")
		}
	}
}
