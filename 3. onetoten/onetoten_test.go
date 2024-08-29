package main

import (
	"testing"
)

type Example struct {
	Value		int
	Expected 	bool
}

var Examples = []Example{
	{0, false},
	{1, true},
	{2, true},
	{3, true},
	{4, true},
	{5, true},
	{6, true},
	{7, true},
	{8, true},
	{9, true},
	{10, true},
	{11, false},
}

func TestOnetoTen(t *testing.T) {
	t.Run("betweenOneAndTen works correctly", func(t *testing.T) {
		for _, example := range Examples {
			result := betweenOneAndTen(example.Value)
			if result != example.Expected {
				t.Error(
					"For", example.Value,
					"Expected", example.Expected,
					"Got", result,
				)
			}
		}
	})
}