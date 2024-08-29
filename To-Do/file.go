package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ConvertTodosToJson(todos ...ToDo) (jsonData []byte) {
	var jsonTodos []map[string]interface{}
	for _, val := range todos {
		jsonTodos = append(jsonTodos, map[string]interface{}{ "id": val.Id, "item": val.Item, "status": val.Status })
	}
	jsonData, err := json.Marshal(jsonTodos)
	check(err, "Failed to convert Todos to JSON")
	return
}

func ReadJSON(fileName string) []ToDo {
	var data []ToDo

	file, _ := os.ReadFile(fileName)
	err := json.Unmarshal(file, &data)
	check(err, fmt.Sprintf("Failed to read JSON %s", fileName))
	
	return data
}

func WriteJson(fileName string, data []byte) {
    file, err := os.Create(fileName)
    check(err, fmt.Sprintf("Failed to create JSON %s", fileName))

    defer file.Close()

    writeErr := os.WriteFile(fileName, data, 0644)
    check(writeErr, fmt.Sprintf("Failed to write JSON %s", fileName))
}
