package main

import (
	"net/http"
	"types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
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
	router.GET("api/todo/:id", getTodo)
	router.POST("api/todo/create", postTodo)
	router.DELETE("api/todo/remove", deleteTodo)
	// router.GET("/todo/create")

	router.Run("localhost:8001")
}

// getTodoList: The handler to response the TODO list
func getTodoList(c *gin.Context) {

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, myTodoList)
}

// getTodo: The handler for response the TODO by Id
func getTodo(c *gin.Context) {
	id := c.Param("id") // Get the value from api/todo/:id

	// Search the matched TODO from the list by Id
	for _, todo := range myTodoList.Todos {
		if todo.Id.String() == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	// If not found, response 204
	c.Writer.WriteHeader(http.StatusNoContent)
}

// postTodo: The handler to add new TODO to TODO list
func postTodo(c *gin.Context) {
	var newTodo types.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		// return
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
	newTodo.Id = uuid.New() // Set Id
	myTodoList.Todos = append(myTodoList.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// deleteTodo: The handler to delete an exist TODO from TODO list
func deleteTodo(c *gin.Context) {
	var deleteTodo types.Todo
	if err := c.BindJSON(&deleteTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	removeIndex := slices.Index(myTodoList.Todos, deleteTodo) // We can find the index if the request contains the full values of the item.

	// But if we only has ID from the request, then get the index by comparing the Id.
	if removeIndex < 0 {
		for index, todo := range myTodoList.Todos {
			if todo.Id == deleteTodo.Id {
				removeIndex = index
				break
			}
		}
	}

	// Try removing the TODO from list
	if removeIndex >= 0 {
		myTodoList.Todos = slices.Delete(myTodoList.Todos, removeIndex, removeIndex+1)
		// fmt.Println(myTodoList)
	} else {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	}
}
