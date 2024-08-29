package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// var fileName string = "ToDo.json"

type ToDoStatus string

const (
	ToDoStatusNotStarted ToDoStatus = "NOTSTARTED"
	ToDoStatusInProgress ToDoStatus = "INPROGRESS"
	ToDoStatusCompleted  ToDoStatus = "COMPLETED"
)

type Status interface {
	GetStatus() string
}

type ToDo struct {
	Id 			uuid.UUID		`json:"id"`
	Item 		string			`json:"item"`
	Status 		ToDoStatus		`json:"status"`
}

func (todo *ToDo) GetStatus() string {
	switch (todo.Status) {
		case ToDoStatusNotStarted:
			return "Not Started \u274c"
		case ToDoStatusInProgress:
			return "In Progress \u25D4"
		case ToDoStatusCompleted:
			return "Completed \u2705"
		default:
			return "Unknown \u2753"
	}
}

func GetTodoItems() []byte {
	var todos []ToDo
	rows, err := db.Query("SELECT * FROM todo")
	check(err, "Failed to get Todos from DB")
	defer rows.Close()
	for rows.Next() {
        var todo ToDo
		err := rows.Scan(&todo.Id, &todo.Item, &todo.Status)
		check(err, fmt.Sprintf("Failed to read Todos from DB %v", todo))
        todos = append(todos, todo)
    }
	check(rows.Err(), "Failed reading todos from DB")
	js, err := json.Marshal(todos)
	check(err, "Failed converting todos to JSON")
	return js
} 

func GetTodoItem(id string) []byte {
    var todo ToDo

	fmt.Println(id)

    row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	err := row.Scan(&todo.Id, &todo.Item, &todo.Status)
	check(err, fmt.Sprintf("Failed reading todo from DB with id = %s", id))
	js, err := json.Marshal(todo)
	check(err, fmt.Sprintf("Failed converting todo to JSON with id = %s", id))

	// todos := ReadJSON(fileName)
	// for i, val := range todos {
	// 	uuid, err := uuid.Parse(id)
	// 	check(err)
	// 	if (val.Id == uuid) {
	// 		return []byte(fmt.Sprintf("%v", todos[i]))
	// 	}
	// }
	return js
}

func CreateTodoItem(item string, status ToDoStatus) uuid.UUID {
	// todos := ReadJSON(fileName)
	// newTodos := append(todos, ToDo{ Id: uuid.New(), Item: item, Status: status })
	// WriteJson(fileName, ConvertTodosToJson(newTodos...))
	var id = uuid.New()
	_, err := db.Exec("INSERT INTO todo (id, item, status) VALUES (?, ?, ?)", id, item, status)
    check(err, fmt.Sprintf("Failed creating todo with id = %s, item = %s, status = %s", id, item, status))
	return id
}

func DeleteTodoItem(id string) {
	// todos := ReadJSON(fileName)
	// for i, val := range todos {
	// 	uuid, err := uuid.Parse(id)
	// 	check(err)
	// 	if (val.Id == uuid) {
	// 		WriteJson(fileName, ConvertTodosToJson(append(todos[:i], todos[i+1:]...)...))
	// 		return
	// 	}
	// }
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
    check(err, fmt.Sprintf("Failed to delete todo with id = %s", id))
}

func UpdateTodoItem(todo ToDo) {
	// todos := ReadJSON(fileName)
	// for i, val := range todos {
	// 	if (val.Id == todo.Id) {
	// 		todos[i] = todo
	// 		WriteJson(fileName, ConvertTodosToJson(todos...))
	// 		return
	// 	}
	// }
	_, err := db.Exec("UPDATE todo SET item = ?, status = ? WHERE id = ?", todo.Item, todo.Status, todo.Id)
    check(err, fmt.Sprintf("Failed update todo %v", todo))
}