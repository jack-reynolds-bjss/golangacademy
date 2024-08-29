package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getTodoItemsHandler() []byte {
	return GetTodoItems()
}

func getTodoItemHandler(id string) []byte {
	return GetTodoItem(id)
}

func createTodoItemHandler(w http.ResponseWriter, req *http.Request) {
	var todo ToDo

  err := json.NewDecoder(req.Body).Decode(&todo)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	fmt.Println(todo)

	CreateTodoItem(todo.Item, todo.Status)
}

func updateTodoItemHandler(w http.ResponseWriter, req *http.Request) {
	var todo ToDo

  err := json.NewDecoder(req.Body).Decode(&todo)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	UpdateTodoItem(todo)
}

func deleteTodoItemHandler(id string) {
	DeleteTodoItem(id)
}

func todoItemsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	
	id := req.URL.Query().Get("id")
	switch (req.Method) {
		case "GET":
			if (id != "") {
				fmt.Printf("Get todo item %s\n", id)
				w.Write(getTodoItemHandler(id))
			} else {
				fmt.Println("Get All todo items...")
				w.Write(getTodoItemsHandler())
			}
		case "POST":
			fmt.Printf("Update todo item %v\n", req.Body)
			updateTodoItemHandler(w, req)
		case "PUT":
			fmt.Printf("Create todo item %v\n", req.Body)
			createTodoItemHandler(w, req)
		case "DELETE":
			if (id != "") {
				fmt.Printf("Delete todo item %s\n", id)
				deleteTodoItemHandler(id)
			} else {
				fmt.Printf("Invalid Id %s\n", id)
				http.Error(w, "Invalid Id", http.StatusBadRequest)
			}
		case "OPTIONS":
			fmt.Println("Pre-flight")
		default:
			fmt.Println("Not found")
			http.Error(w, "Not Found", http.StatusNotFound)
	}
}