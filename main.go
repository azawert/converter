package main

import "fmt"

func main() {
	const UsdToEur = 0.86
	const UsdToRub = 80

	const EurToRub = UsdToRub / UsdToEur
}

func collectUserInput() (int, string, string) {
	var amount int
	var currentCurrency string
	var wantedCurrency string
	fmt.Println("Collecting user input and count")
	fmt.Println("Enter amount of money")
	fmt.Scan(&amount)
	fmt.Println("Enter current currency")
	fmt.Scan(&currentCurrency)
	fmt.Println("Enter wanted currency")
	fmt.Scan(&wantedCurrency)

	return amount, currentCurrency, wantedCurrency
}

func countNewAmount(amount int, currentCurrency string, wantedCurrency string) int {}
