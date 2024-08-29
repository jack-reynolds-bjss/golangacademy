package main

import (
	"fmt"
)

func printArray(arr []int) {
	fmt.Printf("[")
	for i, val := range arr {
		if (i == len(arr) - 1) {
			fmt.Printf("%d", val)
		} else {
			fmt.Printf("%d, ", val)
		}
	}
	fmt.Printf("]\n")
}

func main() {
	arr := []int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}

	fmt.Println("Array in ascending order")
	ascendingArray := FilterAndSortAscending(arr, both)
	printArray(ascendingArray)

	fmt.Println("Array in descending order")
	descendingArray := FilterAndSortDescending(arr, both)
	printArray(descendingArray)

	fmt.Println("Even in ascending order")
	evenAscendingArray := FilterAndSortAscending(arr, even)
	printArray(evenAscendingArray)

	fmt.Println("Even in descending order")
	evenDescendingArray := FilterAndSortDescending(arr, even)
	printArray(evenDescendingArray)

	fmt.Println("Odd in ascending order")
	oddAscendingArray := FilterAndSortAscending(arr, odd)
	printArray(oddAscendingArray)

	fmt.Println("Odd in descending order")
	oddDescendingArray := FilterAndSortDescending(arr, odd)
	printArray(oddDescendingArray)
}