package main

import (
	"fmt"
)

func main() {
	currencies := map[string]float32{
		"JPY": 130.2,
		"EUR": 0.95,
	}
	fmt.Println(currencies)

	var dollarAmount float32
	var currency string

	fmt.Print("Please select an amount: ")

	fmt.Scan(&dollarAmount)
	if dollarAmount == 0 {
		fmt.Println("This is not a valid amount, it's not a float.")
	} else {
		fmt.Print("Please enter currency: ")
		fmt.Scan(&currency)

		rate, isValid := currencies[currency]
		if isValid {
			m := rate * dollarAmount
			fmt.Printf("The amount of target is the %.2f multiplied by the %.2f: %.2f\n", rate, dollarAmount, m)
		} else {
			fmt.Println("The", currency, "is not in the map.")
		}
	}
}
