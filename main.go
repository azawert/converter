package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	operation, errorOperation := getOperation()

	if errorOperation != nil {
		return
	}

	numbers, errorNumber := getNumbers()

	if errorNumber != nil {
		return
	}

	println(calc(operation, numbers))

}

func getOperation() (string, error) {
	var operation string
	var allowedOperations = [4]string{"SUM", "MED", "AVG"}

	fmt.Println("Введите операцию: SUM, MED, AVG")
	fmt.Scan(&operation)

	for _, value := range allowedOperations {
		if value == operation {
			return operation, nil
		}
	}

	return "", errors.New("Operation not allowed")
}

func getNumbers() ([]int, error) {
	var numbers []int
	for {
		var input string

		fmt.Println("Введите любое число или n для прекращения:")
		_, err := fmt.Scan(&input)
		if err != nil {
			return []int{}, err
		}
		if input == "n" {
			break
		}

		var num int
		_, err = fmt.Sscanf(input, "%d", &num)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число или 'n'")
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func calc(operation string, numbers []int) (float64, error) {
	switch operation {
	case "SUM":
		var sum float64
		for _, number := range numbers {
			sum += float64(number)
			return sum, nil
		}
	case "AVG":
		var sum float64
		for _, number := range numbers {
			sum += float64(number) / float64(len(numbers))
			return sum, nil
		}
	case "MED":
		if len(numbers) == 0 {
			return 0, errors.New("для медианы нужен хотя бы один элемент")
		}
		sorted := make([]int, len(numbers))
		copy(sorted, numbers)
		sort.Ints(sorted)

		if len(sorted)%2 == 1 {
			return float64(sorted[len(sorted)/2]), nil
		}
		return float64(sorted[len(sorted)/2-1]+sorted[len(sorted)/2]) / 2.0, nil
	}
	return 0, errors.New("Operation not allowed")

}
