package main

import "fmt"

func main() {
	firstName, middleName, lastName := InputName()

	fmt.Printf("Hello %s %s %s!\n", firstName, middleName, lastName)
}