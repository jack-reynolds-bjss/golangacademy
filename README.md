# golangacademy
A repository for golang academy. This includes the 10 starter activities, Todo HTTP server, and React UI. 

## Server
This either requires mysql with DBUSER and DBPASS exported environment variables or uncomment JSON and comment out db code to write to todo.JSON.

To export variables

```
export DBUSER=root

export DBPASS=<pass>
```

To start the api

```
cd To-Do

go run .
```

### Supported requests

#### Get all todo items
GET localhost:8090/todoItems

returns list of todo items in JSON format
[{ id: "a1a4f4b3-e4ef-4d60-981e-f910ee3e94b1", item: "Todo description", status: "NOTSTARTED" }, { id: "a1a4f4b3-e4ef-4d60-981e-f910ee3e94b1", item: "Todo description", status: "NOTSTARTED" }]

#### Get a single todo item
GET localhost:8090/todoItems?id={id}

returns a single todo item in JSON format
{ id: "a1a4f4b3-e4ef-4d60-981e-f910ee3e94b1", item: "Todo description", status: "NOTSTARTED" }

#### Create a todo item
PUT localhost:8090/todoItems

Body format - { item: "Todo description", status: "NOTSTARTED" }

#### Update a todo item
POST localhost:8090/todoItems - gets a list of all todo items

Body format - { id: "a1a4f4b3-e4ef-4d60-981e-f910ee3e94b1", item: "Todo description", status: "NOTSTARTED" }

#### Delete a todo item
DELETE localhost:8090/todoItems?id={id}

### Future enhancements
- Would have liked to get concurrency working in the HTTP server, there's a simple example commented out of concurrency working with channels
- Also needs more tests

## UI
To start the UI

```
cd To-Do-UI

yarn

yarn start
```




