package main

import (
	"testing"
)

type SumExample struct {
	Value		[]int
	Expected 	int
}

var SumExamples = []SumExample{
	{[]int{1, 2, 3}, 6},
	{[]int{3, 2, 1}, 6},
	{[]int{15, 7, 8}, 30},
	{[]int{6, 45}, 51},
	{[]int{7, 23, 211, 233, 5, 22}, 501},
}

func TestIsOddEven(t *testing.T) {
	for _, example := range SumExamples {
		result := Sum(example.Value)
		if result != example.Expected {
			t.Error(
				"For", example.Value,
				"Expected", example.Expected,
				"Got", result,
			)
		}
	}
}

type DigitsExample struct {
	Value						digits
	ExpectedSingleDigitsSum 	int
	ExpectedDoubleDigitsSum 	int
	ExpectedTrippleDigitsSum 	int
}

var DigitsExamples = []DigitsExample{
	{digits{ singleDigits: []int{1, 2, 3}, doubleDigits: []int{11, 22, 33}, trippleDigits: []int{111, 222, 333} }, 6, 66, 666},
	{digits{ singleDigits: []int{3, 2, 1}, doubleDigits: []int{33, 22, 11}, trippleDigits: []int{333, 222, 111} }, 6, 66, 666},
	{digits{ singleDigits: []int{4, 5, 6}, doubleDigits: []int{54, 23, 65}, trippleDigits: []int{232, 532, 623} }, 15, 142, 1387},
}

func TestSumSingleDigitsExamples(t *testing.T) {
	for _, example := range DigitsExamples {
		result := SumSingleDigits(example.Value)
		if result != example.ExpectedSingleDigitsSum {
			t.Error(
				"For", example.Value,
				"Expected", example.ExpectedSingleDigitsSum,
				"Got", result,
			)
		}
	}
}

func TestSumDoubleDigitsExamples(t *testing.T) {
	for _, example := range DigitsExamples {
		result := SumDoubleeDigits(example.Value)
		if result != example.ExpectedDoubleDigitsSum {
			t.Error(
				"For", example.Value,
				"Expected", example.ExpectedDoubleDigitsSum,
				"Got", result,
			)
		}
	}
}

func TestSumTrippleDigitsExamples(t *testing.T) {
	for _, example := range DigitsExamples {
		result := SumTrippleDigits(example.Value)
		if result != example.ExpectedTrippleDigitsSum {
			t.Error(
				"For", example.Value,
				"Expected", example.ExpectedTrippleDigitsSum,
				"Got", result,
			)
		}
	}
}