package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadCsv(fileName string) (result []Pupil) {
    file, err := os.Open(fileName)
    check(err)
    defer file.Close()
   
    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1
    data, err := reader.ReadAll()
    check(err)

    for i := 1; i < len(data); i++ {
		dob := data[i][3]
		dobSplit := strings.Split(dob, "-")
		year, yearErr := strconv.Atoi(dobSplit[0])
		month, monthErr := strconv.Atoi(dobSplit[1])
		day, dayErr := strconv.Atoi(dobSplit[2])

		if yearErr != nil || monthErr != nil || dayErr != nil {
			panic(err)
		}

        result = append(result, Pupil{
			Name{
				data[i][0],
				data[i][1],
				data[i][2],
			},
			DOB{
				year,
				month,
				day,
			},
		})
	}

    return
}

func WriteCsv(fileName string, data []Pupil, headers []string) {
    file, err := os.Create(fileName)
    check(err)

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write(headers)
    for _, pupil := range data {
        writer.Write([]string{ pupil.Name.FirstName, pupil.Name.MiddleName, pupil.Name.LastName, fmt.Sprintf("%d-%d-%d", pupil.DOB.Year, pupil.DOB.Month, pupil.DOB.Day)  })
    }
}

func main() {
	pupils := []Pupil{ 
		{ Name{ "Neeltje", "Ljuben", "Friis" }, DOB{ 2005, 5, 23 } },
		{ Name{ "Royal", "Fauna", "Filep" }, DOB{ 2005, 3, 4 } },
		{ Name{ "Porfirio", "Marta", "Strom" }, DOB{ 2004, 12, 7 } },
		{ Name{ "Amare", "Vaso", "Donoghue" }, DOB{ 2005, 1, 14 } },
		{ Name{ "Irma", "Hilario", "Pasternack" }, DOB{ 2004, 9, 29 } },
	}
	headers := []string{ "firstName", "middleName", "lastName", "dob" }
	fileName := "pupils.csv"

	WriteCsv(fileName, pupils, headers)

	result := ReadCsv(fileName)

	for _, pupil := range result {
		fmt.Printf("-----\nfull-name : %s\nmiddle-name : %s\nsurname : %s\ndob : %s\nage : %d\n", pupil.FullName(), pupil.Name.MiddleName, pupil.Name.LastName, fmt.Sprintf("%d-%d-%d", pupil.DOB.Year, pupil.DOB.Month, pupil.DOB.Day), pupil.Age())
	}
}