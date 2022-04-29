package main

import (
	"net/http"
	"types"

	"github.com/gin-gonic/gin"
)

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

	// Init Gin router
	router := gin.Default()
	router.GET("api/todo", getTodoList)
	// router.GET("/todo/create")

	router.Run("localhost:8001")
}

func getTodoList(c *gin.Context) {

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, myTodoList)
}
