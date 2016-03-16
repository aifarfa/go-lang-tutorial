package main

import (
	"fmt"
	// "log"
)

func main() {
	// var numbers []int = []int{1, 2, 3}
	numbers := sampleSlice()
	logInfo(numbers)

	numbers = append(numbers, 1, 2, 3, 4, 5) // append values
	logInfo(numbers)

	pop1 := numbers[0 : len(numbers)-1] // remove last
	logInfo(pop1)

	pop2 := numbers[1 : len(numbers)] // remove first
	logInfo(pop2)
}

func sampleSlice() []int {
 	 return make([]int, 0, 5)
}

func logInfo(slice []int) {
	pointer := &slice // address of slice
	fmt.Printf("%p: %v, length: %d, capacity: %d\n", pointer, pointer, len(*pointer), cap(*pointer))
}
