package main

import (
	"errors"
	"fmt"
)

const usdToEuro = 0.86
const usdToRub = 78.25
const euroToRub = usdToRub / usdToEuro

func main() {
	user()
	var amount float64
	var valutaIn, valutaTo string

	fmt.Print("Выберите исходнкю валюту для конвертации: USD, EUR, RUB: ")
	fmt.Scan(&valutaIn)
	ishodValuta, err := proverkaValuti(valutaIn)
	for {
		if err != nil {
			fmt.Println("Неправильно указана исходная валюта: ", valutaIn)
			break
		}

		fmt.Println("Введите сумму для конвертации: ")
		fmt.Scan(&amount)
		summa, err := proverkaAmount(amount)
		if err != nil {
			fmt.Println("Неправильно указана сумма: ", summa)
			break
		}

		valutaVibor1, valutaVibor2 := viborValuti(ishodValuta)
		fmt.Println("Выберите целевую валюту для конвертации: ", valutaVibor1, "/", valutaVibor2)
		fmt.Scan(&valutaTo)
		vibor, err := proverkaValuti(valutaTo)
		if err != nil {
			fmt.Println("Неправильно указана целевая валюта: ", vibor)
			break
		}

		fmt.Printf("Итого: %.2f", convert(amount, valutaIn, valutaTo))
		break
	}
}

func user() string {
	var name string
	fmt.Printf("Здравствуйте! Представьтесь пожалуйста: ")
	fmt.Scan(&name)
	return name
}

func convert(amount float64, from string, to string) float64 {
	switch {
	case from == "USD" && to == "EUR":
		return amount * usdToEuro
	case from == "EUR" && to == "USD":
		return amount * (1 / usdToEuro)
	case from == "USD" && to == "RUB":
		return amount * usdToRub
	case from == "RUB" && to == "USD":
		return amount / usdToRub
	case from == "EUR" && to == "RUB":
		return amount * (usdToRub / usdToEuro)
	case from == "RUB" && to == "EUR":
		return amount / euroToRub
	default:
		return 0
	}
}

func viborValuti(valuta string) (string, string) {
	switch valuta {
	case "USD":
		return "EUR", "RUB"
	case "EUR":
		return "USD", "RUB"
	default:
		return "USD", "EUR"
	}
}

func proverkaValuti(valuta string) (string, error) {
	if valuta != "USD" && valuta != "EUR" && valuta != "RUB" {
		return "", errors.New("no_params_error")
	} else {
		return valuta, nil
	}
}

func proverkaAmount(amount float64) (float64, error) {
	if amount < 0 {
		return amount, errors.New("no_params_error")
	}
	return amount, nil
}
