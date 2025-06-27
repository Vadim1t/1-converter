package main

import "fmt"

func main() {
	user()
	const usdToEuro = 0.86
	const usdToRub = 78.25
	const euroToRub = (1 / usdToEuro) * usdToRub
	fmt.Print(euroToRub)
}

func user() string {
	var name string
	fmt.Printf("Здравствуйте! Представьтесь пожалуйста: ")
	fmt.Scan(&name)
	return name
}

func convert(amount float64, from string, to string) float64 {
	return 0
}
