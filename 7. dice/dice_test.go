package main

import (
	"fmt"
	"testing"
)

type DiceExample struct {
	Roll1		int
	Roll2		int
	Expected 	string
}

var DiceExamples = []DiceExample{
	{1, 1, "SNAKE-EYES-CRAPS"},
	{1, 2, "LOSS-CRAPS"},
	{1, 3, "NEUTRAL"},
	{1, 4, "NEUTRAL"},
	{1, 5, "NEUTRAL"},
	{1, 6, "NATURAL"},
	{2, 1, "LOSS-CRAPS"},
	{2, 2, "NEUTRAL"},
	{2, 3, "NEUTRAL"},
	{2, 4, "NEUTRAL"},
	{2, 5, "NATURAL"},
	{2, 6, "NEUTRAL"},
	{3, 1, "NEUTRAL"},
	{3, 2, "NEUTRAL"},
	{3, 3, "NEUTRAL"},
	{3, 4, "NATURAL"},
	{3, 5, "NEUTRAL"},
	{3, 6, "NEUTRAL"},
	{4, 1, "NEUTRAL"},
	{4, 2, "NEUTRAL"},
	{4, 3, "NATURAL"},
	{4, 4, "NEUTRAL"},
	{4, 5, "NEUTRAL"},
	{4, 6, "NEUTRAL"},
	{5, 1, "NEUTRAL"},
	{5, 2, "NATURAL"},
	{5, 3, "NEUTRAL"},
	{5, 4, "NEUTRAL"},
	{5, 5, "NEUTRAL"},
	{5, 6, "NATURAL"},
	{6, 1, "NATURAL"},
	{6, 2, "NEUTRAL"},
	{6, 3, "NEUTRAL"},
	{6, 4, "NEUTRAL"},
	{6, 5, "NATURAL"},
	{6, 6, "LOSS-CRAPS"},
}

func TestDiceResult(t *testing.T) {
	for _, example := range DiceExamples {
		result := DiceResult(example.Roll1, example.Roll2)
		if result != example.Expected {
			t.Error(
				"For", fmt.Sprintf("%d-%d", example.Roll1, example.Roll2),
				"Expected", example.Expected,
				"Got", result,
			)
		}
	}
}

func TestRollDice(t *testing.T) {
	result := RollDice()
	if result != "SNAKE-EYES-CRAPS" && result != "LOSS-CRAPS" && result != "NATURAL" && result != "NEUTRAL" {
		t.Error(
			"For", "RollDice",
			"Expected", "SNAKE-EYES-CRAPS or LOSS-CRAPS or NATURAL or NEUTRAL",
			"Got", result,
		)
	}
}