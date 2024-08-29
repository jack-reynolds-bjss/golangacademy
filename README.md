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




