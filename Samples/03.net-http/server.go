package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"types"
)

var counter int
var myTodoList types.TodoPageData

const JSON_FILE_PATH = "./files/todo-list.json"

func main() {
	// Read json file as bytes
	bytes := readJsonFile(JSON_FILE_PATH)

	// Deserialize to struct
	var todos []types.Todo
	json.Unmarshal([]byte(bytes), &todos)

	// Debug
	// todosJson, _ := json.MarshalIndent(todos, "", "\t")
	// log.Println(string(todosJson))

	myTodoList = types.TodoPageData{
		PageTitle: "My TODO list",
		Todos:     todos,
		// Todos: []types.Todo{
		// 	{Title: "Task A", IsDone: false},
		// 	{Title: "Task B", IsDone: true},
		// 	{Title: "Task C", IsDone: true},
		// },
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCounter)  // "/count"
	http.HandleFunc("/count/", handlerCounter) // "/count/xxxx"

	http.HandleFunc("/todo", handlerTodoList)          // "/todo"
	http.HandleFunc("/todo/create", handlerTodoCreate) // "/todo/create"
	// http.HandleFunc("/todo/save", handlerTodoSave)     // "todo/save"

	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)

	// Use http.ResponseWriter
	// rw.Write([]byte(`My website`))

	// Or use fmt.Fprintf
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
		tmpl := template.Must(template.New("todo-create.html").ParseFiles("./todo-create.html", "./header.html"))
		tmpl.Execute(rw, nil)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(rw, "PostFrom = %v\n", req.PostForm)
		todo := req.FormValue("todo")           // "todo" is the name of the input dom
		isDone := req.FormValue("isDone") != "" // "isDone" is the name of the checkbox dom
		// fmt.Fprintf(rw, "Todo = %s, IsDone = %v\n", todo, isDone)

		// myTodoList.Todos = append(myTodoList.Todos, types.Todo{Title: todo, IsDone: isDone})
		todosNew := &(myTodoList.Todos)
		*todosNew = append(*todosNew, types.Todo{Title: todo, IsDone: isDone})

		// Write to json file
		todosJson, _ := json.MarshalIndent(todosNew, "", "\t")
		writeJsonFile(JSON_FILE_PATH, string(todosJson))

		http.Redirect(rw, req, "/todo", http.StatusSeeOther)
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

func handlerTodoList(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		tmpl := template.Must(template.New("todo-list.html").Funcs(template.FuncMap{"inc": inc}).ParseFiles("./header.html", "./todo-list.html"))
		// tmpl := template.Must(template.Must(template.ParseFiles("./todo-list.html")).ParseFiles("./header.html"))
		tmpl.Execute(rw, myTodoList)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}

		// (Optional) To get URL parameter
		// removeIndexs, isOk := req.URL.Query()["removeId"]
		// if !isOk || len(removeIndexs) < 1 {
		// 	rw.WriteHeader(http.StatusBadRequest)
		// }
		// removeIndex, _ := strconv.Atoi(removeIndexs[0])

		// 2. Using Form post
		removeIndex, _ := strconv.Atoi(req.FormValue("removeId")) // Convert string to int

		todosNew := &(myTodoList.Todos)
		*todosNew = append((*todosNew)[:removeIndex], (*todosNew)[removeIndex+1:]...) // Append elements before and after the removed index.

		// Write to json file
		todosJson, _ := json.MarshalIndent(todosNew, "", "\t")
		writeJsonFile(JSON_FILE_PATH, string(todosJson))

		// Redirect
		http.Redirect(rw, req, "/todo", http.StatusSeeOther)
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

// inc: increment by 1
func inc(i int) int {
	return i + 1
}

// readJsonFile: read json file as bytes
func readJsonFile(path string) []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// writeJsonFile: write json file
func writeJsonFile(path string, str string) {
	jsonFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755) // Set the overwrite permission
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	jsonFile.Sync()
}
