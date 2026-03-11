package main

import "fmt"

func userInput() string {
	var input string
	fmt.Scan(&input)

	return input
}

func calculateCurrency(value float64, sourceCurrency string, targetCurrency string) {

}

func main() {
	const EUR = 0.8607
	const RUB = 79.15
	var countRub = 1000.0
	var countEur = 150.0
	resultEur := EUR * (countRub / RUB)
	resultRub := RUB * (countEur / EUR)
	fmt.Printf("Конвертация из RUB в EUR = %F\n", resultEur)
	fmt.Printf("Конвертация из EUR в RUB = %F\n", resultRub)
}
