package main

import "fmt"

func main() {
	const UsdToEur = 0.86
	const UsdToRub = 80

	const EurToRub = UsdToRub / UsdToEur
}

func collectUserInputAndCount(amount int, currentCurrency string, wantedCurrency string) int {
	fmt.Println("Collecting user input and count")
	fmt.Println("Enter amount of money")
	fmt.Scan(&amount)
	fmt.Println("Enter current currency")
	fmt.Scan(&currentCurrency)
	fmt.Println("Enter wanted currency")
	fmt.Scan(&wantedCurrency)
}
