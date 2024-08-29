package main

import "fmt"

func main() {
	cities := [][]string{ { "Abu Dhabi" }, { "London" }, { "Washington D.C." }, { "Montevideo" }, { "Vatican City" }, { "Caracas" }, { "Hanoi" }}
	headers := []string{ "name" }
	fileName := "cities.csv"

	WriteCsv(fileName, cities, headers)

	readCities := ReadCsv(fileName)
	SortArray(readCities)

	fmt.Println(readCities)
}
