package main

import (
	"fmt"
	"math/rand"
	"time"
)

const chances = 5

func main() {
	// sampleLoop()
	x := getRandomValue()

	var input int
	for counter := 0; counter < chances; counter++ {
		fmt.Scanln(&input)
		hint(x, input)

		if input == x {
			break
		}
		remainingChances(counter)
	}
}

func hint(expect, actual int) {
	switch {

	case actual > expect:
		fmt.Println(actual, "greater than expect")

	case actual < expect:
		fmt.Println(actual, "is less than expect")

	default:
		fmt.Println("Correct!")
		break
	}
}

func remainingChances(used int) {
	remain := chances - used - 1

	if remain == 0 {
		fmt.Println("Game over..")
	} else {
		fmt.Println("remaining chances:", remain)
	}
}

func getRandomValue() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(100) + 1
}

func sampleLoop() {
	// i := 0
	for i := 0; i < 8; i++ {
		fmt.Println("go:", i)
		// i++
	}
}
