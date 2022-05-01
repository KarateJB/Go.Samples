package main

import (
	"net/http"
	"types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

var myTodoList types.TodoPageData

// @Title TODO API
// @Version 1.0
// @Description TODO API sample by Gin
// @Accept json
// @Produce json
// @Host localhost:8001
func main() {
	myTodoList = types.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []types.Todo{
			{Id: uuid.MustParse("aa3cdd2f-17b9-4f43-9eb0-af56b42908c5"), Title: "Task A", IsDone: false},
			{Id: uuid.MustParse("bbf5d05c-c442-4869-8326-ab5cfa832f6a"), Title: "Task B", IsDone: true},
			{Id: uuid.MustParse("cca89c32-a0d9-43c9-84e2-ae1224c5d755"), Title: "Task C", IsDone: true},
		},
	}

	// Init Gin router
	router := gin.Default()
	router.GET("api/todo", getTodoList)
	router.GET("api/todo/:id", getTodo)
	router.POST("api/todo", postTodo)
	router.PUT("api/todo", putTodo)
	router.DELETE("api/todo", deleteTodo)

	router.Run("localhost:8001")
}

// @Title Get TODO list
// @Version 1.0
// @Description The handler to response the TODO list
// @Router /api/todo
// @Success 200 {types.TodoPageData}
// getTodoList: The handler to response the TODO list
func getTodoList(c *gin.Context) {

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, myTodoList)
}

// @Title Get a TODO by its Id
// @Version 1.0
// @Description The handler for response the TODO by Id
// @Router /api/todo/:id [get]
// @Success 200 {types.Todo}
// @Success 204
// getTodo: The handler for response the TODO by Id
func getTodo(c *gin.Context) {
	id := c.Param("id") // Get the value from api/todo/:id

	// Search the matched TODO from the list by Id.
	for _, todo := range myTodoList.Todos {
		if todo.Id.String() == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	// If not found, response 204
	c.Writer.WriteHeader(http.StatusNoContent)
}

// @Title Create a new TODO
// @Version 1.0
// @Description The handler to add new TODO to TODO list
// @Router /api/todo [post]
// @Success 201 {types.Todo}
// @Failure 401
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

// @Title Edit a TODO
// @Version 1.0
// @Description The handler to edit a TODO
// @Router /api/todo [put]
// @Success 200
// @Failure 401
// putTodo: The handler to edit a TODO
func putTodo(c *gin.Context) {
	var editTodo types.Todo
	if err := c.BindJSON(&editTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	for index, todo := range myTodoList.Todos {
		if todo.Id == editTodo.Id {
			myTodoList.Todos[index].Title, myTodoList.Todos[index].IsDone = editTodo.Title, editTodo.IsDone

			// range copies the values from the slice that iterated over, so below code wont work.
			// todo.Title = editTodo.Title
			// todo.IsDone = editTodo.IsDone

			return
		}
	}

	c.Writer.WriteHeader(http.StatusBadRequest)
}

// @Title Delete a TODO
// @Version 1.0
// @Description The handler to delete an exist TODO from TODO list
// @Router /api/todo [delete]
// @Success 200
// @Failure 401
// @Failure 422
// deleteTodo: The handler to delete an exist TODO from TODO list
func deleteTodo(c *gin.Context) {
	var deleteTodo types.Todo
	if err := c.BindJSON(&deleteTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	// We can find the index if the request contains the full values of the item.
	removeIndex := slices.Index(myTodoList.Todos, deleteTodo)

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
