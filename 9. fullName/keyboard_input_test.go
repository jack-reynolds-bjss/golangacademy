package main

import (
	"testing"
)

type PersonExample struct {
	FirstName string
	MiddleName string
	LastName string
	Expected Person
}

var PeopleExamples = []PersonExample{
	{"Gandalf", "the", "Grey", Person{ "Gandalf the Grey", Name{ "Gandalf", "the", "Grey" } }},
	{"Aragorn", "son of", "Arathorn", Person{ "Aragorn son of Arathorn", Name{ "Aragorn", "son of", "Arathorn" } }},
	{"Legolas", "son of", "Thranduil", Person{ "Legolas son of Thranduil", Name{ "Legolas", "son of", "Thranduil" } }},
	{"Gimli", "son of", "Gloin", Person{ "Gimli son of Gloin", Name{ "Gimli", "son of", "Gloin" } }},
}

func TestInputName(t *testing.T) {
	for _, person := range PeopleExamples {
		result := InputName()
		if result != person.Expected {
			t.Error(
				"For", person.Expected.FullName,
				"Expected", person.Expected,
				"Got", result,
			)
		}
	}
}