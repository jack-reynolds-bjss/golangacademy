package main

import "fmt"

func main() {
	var i int
	var response string = ""
	var icon string = ""

	fmt.Print("Enter a number: ")
	fmt.Scan(&i)

	if (!betweenOneAndTen(i)) {
		response = " not"
		icon = "\u274c "
	} else {
		icon = "\u2705 "
	}

	fmt.Printf("%sThe number you have entered is%s between one and ten\n", icon, response)
}