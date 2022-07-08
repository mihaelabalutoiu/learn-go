package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards <= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else if isHeistOn == false {
		fmt.Println("Plan a better disguise next time?")
	}

	fmt.Println(isHeistOn)
	openedVault := rand.Intn(100)

	if isHeistOn == true && openedVault <= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == false {
		fmt.Println("The vault can't be opened")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("It's failed!")
		case 1:
			isHeistOn = true
			fmt.Println("Turns out vault doors don't open from the inside")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println(amtStolen)
	}
}
