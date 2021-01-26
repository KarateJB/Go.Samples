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
var myTodoList types.TodoPageData

func main() {
	myTodoList = types.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []types.Todo{
			{Title: "Task A", IsDone: false},
			{Title: "Task B", IsDone: true},
			{Title: "Task C", IsDone: true},
		},
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCounter)  // "/count"
	http.HandleFunc("/count/", handlerCounter) // "/count/xxxx"

	http.HandleFunc("/todo", handlerTodoList)          // "/todo"
	http.HandleFunc("/todo/create", handlerTodoCreate) // "/todo/create"

	log.Fatal(http.ListenAndServe("localhost:8001", nil))
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

func handlerTodoCreate(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		tmpl := template.Must(template.Must(template.ParseFiles("./todo-create.html")).ParseFiles("./header.html"))
		tmpl.Execute(rw, nil)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(rw, "PostFrom = %v\n", req.PostForm)
		todo := req.FormValue("todo")
		isDone := req.FormValue("isDone") != ""
		// fmt.Fprintf(rw, "Todo = %s, IsDone = %v\n", todo, isDone)
		myTodoList.Todos = append(myTodoList.Todos, types.Todo{Title: todo, IsDone: isDone})
		http.Redirect(rw, req, "/todo", http.StatusPermanentRedirect)
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

func handlerTodoList(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if len(myTodoList.Todos) == 3 {
			tmpl := template.Must(template.New("todo-list.html").Funcs(template.FuncMap{"inc": inc}).ParseFiles("./header.html", "./todo-list.html"))
			// tmpl := template.Must(template.Must(template.ParseFiles("./todo-list.html")).ParseFiles("./header.html"))
			tmpl.Execute(rw, myTodoList)
		} else {
			fmt.Fprintf(rw, "RemoveId = %v", myTodoList.Todos)
		}

	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}
		removeIndex, _ := strconv.Atoi(req.FormValue("removeId"))
		// todosNew := myTodoList.Todos
		nextIndex := removeIndex + 1
		myTodoList.Todos = append(myTodoList.Todos[:removeIndex], myTodoList.Todos[nextIndex:]...)
		fmt.Fprintf(rw, "RemoveId = %v", myTodoList.Todos)
		// http.Redirect(rw, req, "/todo", http.StatusPermanentRedirect)
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

// inc: increment by 1
func inc(i int) int {
	return i + 1
}
