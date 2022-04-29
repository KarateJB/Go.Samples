package main

import (
	"net/http"
	"types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var myTodoList types.TodoPageData

func main() {
	myTodoList = types.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []types.Todo{
			{Id: uuid.New(), Title: "Task A", IsDone: false},
			{Id: uuid.New(), Title: "Task B", IsDone: true},
			{Id: uuid.New(), Title: "Task C", IsDone: true},
		},
	}

	// Init Gin router
	router := gin.Default()
	router.GET("api/todo", getTodoList)
	router.POST("api/todo/create", postTodoList)
	// router.GET("/todo/create")

	router.Run("localhost:8001")
}

// getTodoList: The handler to response the TODO list
func getTodoList(c *gin.Context) {

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, myTodoList)
}

// postTodoList: The handler to add new TODO to TODO list
func postTodoList(c *gin.Context) {
	var newTodo types.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		// log.Fatal(err)
		return
	}
	newTodo.Id = uuid.New() // Set Id
	myTodoList.Todos = append(myTodoList.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
