package main

import "fmt"

func main() {
	const usdToEuro = 0.86
	const usdToRub = 78.25
	const euroToRub = (1 / usdToEuro) * usdToRub
	fmt.Print(euroToRub)
}
