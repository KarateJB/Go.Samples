package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"types"
)

var counter int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCounter)  // "/count"
	http.HandleFunc("/count/", handlerCounter) // "/count/xxxx"

	http.HandleFunc("/todo", handlerTodo) // "/todo"

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	// rw.Write([]byte(`My website`))
	fmt.Fprintf(rw, "Welcome, your current URL.Path = %q\n", req.URL.Path)
}

func handlerCounter(rw http.ResponseWriter, req *http.Request) {
	counter++
	rw.WriteHeader(http.StatusOK)
	// fmt.Fprintf(rw, "Counter = %d\n", counter)
	html := `<!DOCTYPE html>
			<html>
			<head><title>My counter</title></head>
			<body><h2>Counter = ` + strconv.Itoa(counter) + `</h2></body>
			</html>`
	rw.Write([]byte(html))
}

func handlerTodo(rw http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.Must(template.ParseFiles("./layout.html")).ParseFiles("./header.html"))
	data := types.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []types.Todo{
			{Title: "Task A", IsDone: false},
			{Title: "Task B", IsDone: true},
			{Title: "Task C", IsDone: true},
		},
	}
	tmpl.Execute(rw, data)
}
