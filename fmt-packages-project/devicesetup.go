package main

import (
	"fmt"
)

func main() {
	fmt.Println("What`s your name?")
	var name string
	var age int

	fmt.Scan(&name)
	fmt.Println("How old are you?")
	fmt.Scan(&age)

	newMessage := fmt.Sprintf("Hi %s, nice to meet you!", name)
	fmt.Println(newMessage)
}
