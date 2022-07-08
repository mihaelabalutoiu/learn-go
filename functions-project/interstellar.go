package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("You have left", fuel, "of the fuel")
}

func calculateFuel(planet string) int {
	var fuel int
	if planet == "Venus" {
		fuel = 300000
	} else if planet == "Mercury" {
		fuel = 500000
	} else if planet == "Mars" {
		fuel = 700000
	} else {
		fmt.Println(0)
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome everyone to the planet:", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	fuel := 1000000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
	calculateFuel("Venus")
}
