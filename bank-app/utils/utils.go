package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE string = "balance.txt"

func balanceToFloat(text []byte) float64 {
	balanceText := string(text)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		log.Fatal("error converting streing to float " + err.Error())
	}
	return balance
}

func WriteBalanceToFile(balance float64) {
	balanceToString := fmt.Sprint(balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceToString), 0644)
}

func ReadBalanceFromFile() float64 {
	data, err := os.ReadFile(ACCOUNT_BALANCE_FILE)
	if err != nil {
		log.Fatal("error reading file. " + err.Error())
	}
	balance := balanceToFloat(data)
	return balance
}
