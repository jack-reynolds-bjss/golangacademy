package main

import (
	"fmt"
)

func main() {
	var day int
	var month int
	var year int

	fmt.Print("Enter day: ")
	fmt.Scan(&day)
	fmt.Print("Enter month: ")
	fmt.Scan(&month)
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	age := CalculateAge(year, month, day)

	fmt.Printf("Your age is: %d\n", age)
}