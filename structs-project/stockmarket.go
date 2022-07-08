package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

type Stock struct {
	name  string
	price float32
}

func (s *Stock) updateStock() {
	change := randomNumberGen(-10.0, 10.0)
	s.price += change
}

func displayMarket(market []Stock) {
	for _, element := range market {
		fmt.Println(element)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	stockMarket := []Stock{{"GOOG", 2313.50}, {"AAPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.00}}

	displayMarket(stockMarket)
	stockMarket[1].updateStock()
	fmt.Println("===")

	displayMarket(stockMarket)
}
