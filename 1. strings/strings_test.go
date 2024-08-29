package main

import (
	"testing"
)

func TestStrings(t *testing.T) {
	t.Run("Strings work correctly", func(t *testing.T) {
		string1, string2, string3 := Strings()
		if string1 != "Hello" {
			t.Error(
				"For", "string1",
				"Expected", "Hello",
				"Got", string1,
			)
		}
		if string2 != "go"{
			t.Error(
				"For", "string2",
				"Expected", "go",
				"Got", string2,
			)
		}
		if string3 != "academy" {
			t.Error(
				"For", "string3",
				"Expected", "academy",
				"Got", string3,
			)
		}
	})
}