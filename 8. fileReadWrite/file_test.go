package main

import (
	"testing"
)

type Example struct {
	Value		[]string
	Expected 	[]string
}

var Examples = []Example{
	{[]string{ "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi" }, []string{ "Abu Dhabi", "Caracas", "Hanoi", "London", "Montevideo", "Vatican City", "Washington D.C." }},
}

func TestDiceResult(t *testing.T) {
	for _, example := range Examples {
		result := example.Value
		SortArray(result)
		for i, val := range result {
			if val != example.Expected[i] {
				t.Error(
					"For", example.Value[i],
					"Expected", example.Expected[i],
					"Got", val,
				)
			}
		}
	}
}