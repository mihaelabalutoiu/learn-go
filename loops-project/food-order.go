package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	fmt.Scan(&input)
	return input
}

func contains(menu []string, order string) bool {
	for _, i := range menu {
		if order == i {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string

	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 4
		} else if order != "quit" {
			fmt.Println("This item is not on the menu")
		}
	}

	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("The remaining amount: ", amount, "need to be paid to the cashier.")
	}
	fmt.Println("The total for the order is", total)
}
