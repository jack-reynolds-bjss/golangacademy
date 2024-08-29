package main

import (
	"testing"
)

type Person struct {
	FirstName	string
	MiddleName 	string
	LastName  	string
}

var PeopleExamples = []Person{
	{"Gandalf", "the", "Grey"},
	{"Aragorn", "son of", "Arathorn"},
	{"Legolas", "son of", "Thranduil"},
	{"Gimli", "son of", "Gloin"},
}

func TestInputName(t *testing.T) {

	t.Run("InputName works correct", func(t *testing.T) {
		for _, person := range PeopleExamples {
			firstName, middleName, lastName := InputName()
			if firstName != person.FirstName {
				t.Error(
					"For", person.FirstName,
					"Expected", person.FirstName,
					"Got", firstName,
				)
			}
			if middleName != person.MiddleName {
				t.Error(
					"For", person.MiddleName,
					"Expected", person.MiddleName,
					"Got", middleName,
				)
			}
			if lastName != person.LastName {
				t.Error(
					"For", person.LastName,
					"Expected", person.LastName,
					"Got", lastName,
				)
			}
		}
	})

}