package main

import (
	"fmt"
)

func main() {
	person := InputName()

	fmt.Printf("full-name : %s\n", person.FullName)
	fmt.Printf("middle-name : %s\n", person.Name.MiddleName)
	fmt.Printf("surname : %s\n", person.Name.LastName)
}