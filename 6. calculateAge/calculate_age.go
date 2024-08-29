package main

import (
	"time"

	age "github.com/bearbin/go-age"
)

func CalculateAge(year int, month int, day int) int {
	var dob time.Time = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return age.Age(dob)
}