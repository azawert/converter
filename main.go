package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func main() {
	// Цикл для повторного ввода при ошибке
	for {
		amount, currentCurrency, wantedCurrency, inputError := collectUserInput()

		if inputError != nil {
			fmt.Printf("❌ Error: %v\n\n", inputError)
			continue // Повторяем ввод
		}

		newAmount := countNewAmount(amount, currentCurrency, wantedCurrency)
		fmt.Printf("\n✓ Result: %.2f %s = %.2f %s\n", float64(amount), currentCurrency, newAmount, wantedCurrency)
		break // Успешно, выходим из цикла
	}
}

func collectUserInput() (int, string, string, error) {
	validCurrencies := []string{"EUR", "USD", "RUB"}

	// Ввод и валидация текущей валюты
	fmt.Printf("Enter a valid currency [%s]: ", strings.Join(validCurrencies, ", "))
	currentCurrency, err := inputAndValidateCurrency(validCurrencies)
	if err != nil {
		return 0, "", "", err
	}

	// Ввод и валидация суммы
	fmt.Print("Enter a value: ")
	amount, err := inputAndValidateNumber()
	if err != nil {
		return 0, "", "", err
	}

	// Ввод и валидация желаемой валюты
	fmt.Print("Enter a wanted currency: ")
	wantedCurrency, err := inputAndValidateCurrency(validCurrencies)
	if err != nil {
		return 0, "", "", err
	}

	// Проверка, что валюты разные
	if wantedCurrency == currentCurrency {
		return 0, "", "", errors.New("currencies must be different")
	}

	return amount, currentCurrency, wantedCurrency, nil
}

// Функция ввода и проверки валюты
func inputAndValidateCurrency(validCurrencies []string) (string, error) {
	var currency string
	fmt.Scanln(&currency)
	currency = strings.ToUpper(strings.TrimSpace(currency))

	if !validateCurrency(currency, validCurrencies) {
		return "", errors.New("invalid currency provided")
	}

	return currency, nil
}

// Функция валидации валюты
func validateCurrency(currency string, validCurrencies []string) bool {
	return slices.Contains(validCurrencies, currency)
}

// Функция ввода и проверки числа
func inputAndValidateNumber() (int, error) {
	var amount int
	_, err := fmt.Scanln(&amount)

	if err != nil {
		// Очистка буфера при ошибке
		var discard string
		fmt.Scanln(&discard)
		return 0, errors.New("invalid amount provided")
	}

	if amount <= 0 {
		return 0, errors.New("amount must be positive")
	}

	return amount, nil
}

func countNewAmount(amount int, currentCurrency string, wantedCurrency string) float64 {
	const UsdToEur = 0.86
	const UsdToRub = 80.0
	const EurToUsd = 1.16
	const RubToUsd = 0.0125

	if currentCurrency == wantedCurrency {
		return float64(amount)
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
		return amountInUsd
	case "EUR":
		return amountInUsd * UsdToEur
	case "RUB":
		return amountInUsd * UsdToRub
	default:
		return 0
	}
}
