package main

import (
	"fmt"
	"testing"
)

type CalculateAgeExample struct {
	Year		int
	Month		int
	Day			int
	Expected 	int
}

var SumExamples = []CalculateAgeExample{
	{1995, 1, 26, 29},
	{1990, 5, 10, 34},
	{2004, 4, 14, 20},
	{2000, 8, 13, 24},
	{2000, 8, 14, 23},
}

func TestCalculateAge(t *testing.T) {
	for _, example := range SumExamples {
		result := CalculateAge(example.Year, example.Month, example.Day)
		if result != example.Expected {
			t.Error(
				"For", fmt.Sprintf("%d-%d-%d", example.Year, example.Month, example.Day),
				"Expected", example.Expected,
				"Got", result,
			)
		}
	}
}