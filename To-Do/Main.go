package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var port int = 8090
// var wg = sync.WaitGroup{}


// func printToDo(todos ...ToDo) {
// 	for _, val := range todos {
// 		fmt.Printf("%s - %s\n", val.item, getStatus(val))
// 	}
// }

// func getTodoitems() []ToDo {
// 	todo1 := ToDo{ item: "Buy eggs", status: ToDoStatusNotStarted }
// 	todo2 := ToDo{ item: "Sleep", status: ToDoStatusInProgress }
// 	todo3 := ToDo{ item: "Watch LOTR", status: ToDoStatusNotStarted }
// 	todo4 := ToDo{ item: "Write a novel", status: ToDoStatusCompleted }
// 	todo5 := ToDo{ item: "Finish ToDo", status: ToDoStatusInProgress }

// 	// printToDo(todo1, todo2, todo3, todo4, todo5)

// 	jsonData := ConvertTodosToJson(todo1, todo2, todo3, todo4, todo5)
// 	WriteJson(fileName, jsonData)

// 	return []ToDo{ todo1, todo2, todo3, todo4, todo5 }
// }

// func queueTodoItem(ch chan<- string, todo ToDo) {
// 	ch <- todo.Item
// 	wg.Done()
// }

// func queueTodoStatus(ch chan<- string, todo ToDo) {
// 	ch <- todo.GetStatus()
// 	wg.Done()
// }

// func raceCondition(todoItems []ToDo) {
// 	ch := make(chan string, len(todoItems) * 2)

// 	for _,val := range todoItems {
// 		wg.Add(1)
// 		go queueTodoItem(ch, val)
// 		wg.Wait()
		
// 		wg.Add(1)
// 		go queueTodoStatus(ch, val)
// 		wg.Wait()
		
// 		fmt.Printf("ToDo: %s - %s\n", <-ch, <-ch)
// 	}
// }

func startDB() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "golangacademy",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")
}

func startServer() {
	http.HandleFunc("/todoItems", todoItemsHandler)

	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	
}

func main() {
	// toDoItems := GetTodoItems()
	// raceCondition(toDoItems)

	// AddTodoItem("New", ToDoStatusNotStarted)

	// toDoItems = GetTodoItems()
	// raceCondition(toDoItems)

	// RemoveTodoItem(toDoItems[0].Id)
	
	// toDoItems = GetTodoItems()
	// raceCondition(toDoItems)

	startDB()
	startServer()
}