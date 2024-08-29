package main

import (
	"math/rand/v2"
)

func DiceResult(roll1 int, roll2 int) (result string) {
	switch (roll1 + roll2) {
		case 2:
			result = "SNAKE-EYES-CRAPS"
		case 3, 12:
			result = "LOSS-CRAPS"
		case 7, 11:
			result = "NATURAL"
		default:
			result = "NEUTRAL"
	}
	return
}

func RollDice() string {
	roll1 := rand.IntN(6) + 1
	roll2 := rand.IntN(6) + 1

	return DiceResult(roll1, roll2)
}
