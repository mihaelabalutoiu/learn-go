package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	guestlist := [3]string{"Famous people", "People you know", "Goofy names"}
	catstorage := [3]string{"A Toy Chest", "A Crate", "A Box"}

	fmt.Println("Guestlist:", guestlist)
	fmt.Println("Catstorage:", catstorage)

	TheGuests := getRandomElement(guestlist[:])
	TheLocation := getRandomElement(catstorage[:])
	fmt.Println(TheGuests)
	fmt.Println(TheLocation)
}

func getRandomElement(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	fmt.Println(randomIndex)
	return slice[randomIndex]
}
