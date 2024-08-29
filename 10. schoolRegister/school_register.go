package main

import (
	"fmt"
	"time"

	age "github.com/bearbin/go-age"
)

type Person interface {
	Age() int
	FullName() string
}

type DOB struct {
	Year int
	Month int
	Day int
}

type Name struct {
	FirstName string
	MiddleName string
	LastName string
}

type Pupil struct {
	Name Name
	DOB DOB
}

func (p *Pupil) Age() int {
	var dob time.Time = time.Date(p.DOB.Year, time.Month(p.DOB.Month), p.DOB.Day, 0, 0, 0, 0, time.UTC)

	return age.Age(dob)
}

func (p *Pupil) FullName() string {
	return fmt.Sprintf("%s %s %s", p.Name.FirstName, p.Name.MiddleName, p.Name.LastName)
}