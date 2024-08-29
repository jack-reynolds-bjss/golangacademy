package main

import "fmt"

func InputName() (string, string, string) {
	var firstName string
	var middleName string
	var lastName string

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter middle name: ")
	fmt.Scan(&middleName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	return firstName, middleName, lastName
}