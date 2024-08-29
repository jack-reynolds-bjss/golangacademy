package main

import (
	"fmt"
	"strings"
)

type Name struct {
	FirstName string
	MiddleName string
	LastName string
}

type Person struct {
	FullName string
	Name Name
}

func InputName() (Person) {
	var firstName string
	var middleName string
	var lastName string

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter middle name: ")
	fmt.Scan(&middleName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fullName := fmt.Sprintf("%s %s %s", firstName, middleName, lastName);

	fullNameSplit := strings.Split(fullName, " ")

	person := Person{
		fullName,
		Name{
			fullNameSplit[0],
			fullNameSplit[1],
			fullNameSplit[2],
		},
	}

	return person
}