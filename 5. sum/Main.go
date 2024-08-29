package main

import (
	"fmt"
)

func main() {
	numbers := digits{
		singleDigits: 	[]int{1, 2, 3},
		doubleDigits: 	[]int{11, 22, 3},
		trippleDigits: 	[]int{111, 222, 333},
	}

	singleDigitTotal := SumSingleDigits(numbers)
	doubleDigitTotal := SumDoubleeDigits(numbers)
	trippleDigitTotal := SumTrippleDigits(numbers)
	total := Sum([]int{ singleDigitTotal, doubleDigitTotal, trippleDigitTotal })

	fmt.Printf("Single digits total: %d\n", singleDigitTotal)
	fmt.Printf("Double digits total: %d\n", doubleDigitTotal)
	fmt.Printf("Tripple digits total: %d\n", trippleDigitTotal)
	fmt.Printf("Total: %d\n", total)
}