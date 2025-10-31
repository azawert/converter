package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func main() {

	amount, currentCurrency, wantedCurrency, inputError := collectUserInput()

	if inputError != nil {
		fmt.Println(inputError)
		return
	}

	newAmount := countNewAmount(amount, currentCurrency, wantedCurrency)

	println(newAmount)

}

func collectUserInput() (int, string, string, error) {
	var amount int
	var currentCurrency string
	var wantedCurrency string
	validCurrencies := []string{"EUR", "USD", "RUB"}

	// Ввод текущей валюты
	fmt.Printf("Enter a valid currency %s\n", strings.Join(validCurrencies, ","))
	fmt.Scan(&currentCurrency)
	currentCurrency = strings.ToUpper(currentCurrency)

	if !slices.Contains(validCurrencies, currentCurrency) {
		return 0, "", "", errors.New("Invalid currency provided")
	}

	// ✅ ДОБАВЛЕНО: Очистка буфера после первого ввода
	fmt.Scanln() // Считываем оставшийся символ новой строки

	// Ввод суммы
	fmt.Println("Enter a value")
	_, err := fmt.Scanf("%d", &amount)
	if err != nil {
		return 0, "", "", errors.New("Invalid amount provided")
	}

	// ✅ ДОБАВЛЕНО: Очистка буфера после ввода числа
	fmt.Scanln()

	// Ввод желаемой валюты
	fmt.Println("Enter a wanted currency")
	fmt.Scan(&wantedCurrency)
	wantedCurrency = strings.ToUpper(wantedCurrency)

	if wantedCurrency == currentCurrency {
		return 0, "", "", errors.New("Currencies must be different")
	}

	if !slices.Contains(validCurrencies, wantedCurrency) {
		return 0, "", "", errors.New("Invalid currency provided")
	}

	return amount, currentCurrency, wantedCurrency, nil
}

func countNewAmount(amount int, currentCurrency string, wantedCurrency string) int {
	const UsdToEur = 0.86
	const UsdToRub = 80.0
	const EurToUsd = 1.16
	const RubToUsd = 0.0125

	if currentCurrency == wantedCurrency {
		return int(float64(amount))
	}

	amountInUsd := 0.0
	switch currentCurrency {
	case "USD":
		amountInUsd = float64(amount)
	case "EUR":
		amountInUsd = float64(amount) * EurToUsd
	case "RUB":
		amountInUsd = float64(amount) * RubToUsd
	default:
		return 0
	}

	switch wantedCurrency {
	case "USD":
		return int(amountInUsd)
	case "EUR":
		return int(amountInUsd * UsdToEur)
	case "RUB":
		return int(amountInUsd * UsdToRub)
	default:
		return 0
	}

}
