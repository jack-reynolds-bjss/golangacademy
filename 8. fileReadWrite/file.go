package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func SortArray(cities []string) {
    slices.Sort(cities)
}

func ReadCsv(fileName string) (result []string) {
    file, err := os.Open(fileName)
    check(err)
    defer file.Close()
   
    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1
    data, err := reader.ReadAll()
    check(err)

    for i := 1; i < len(data); i++ {
        result = append(result, data[i][0])
    }

    fmt.Println(result)

    return
}

func WriteCsv(fileName string, data [][]string, headers []string) {
    file, err := os.Create(fileName)
    check(err)

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write(headers)
    for _, row := range data {
        writer.Write(row)
    }
}
