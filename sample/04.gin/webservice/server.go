package main

import (
	"example/webservice/docs"
	"net/http"
	"strconv"
	"strings"
	"types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"golang.org/x/exp/slices"
)

var myTodoList types.TodoPageData

// @Title TODO API
// @Version 1.0
// @Description TODO API sample by Gin
// @Host localhost:8001
// @BasePath /
// @Contact.Name JB
// @Contact.Url https://karatejb.blogspot.com/
// @Contact.Email xxx@demo.com
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
	router.GET("api/todo/:id", getTodo) // The id is required for matching this routing
	// router.GET("api/todo/*id", getTodoById) // The id is optional for matching this routing, e.q. api/todo/ or api/todo/xxx
	router.GET("api/todo/search", searchTodo)
	router.POST("api/todo", postTodo)
	router.PUT("api/todo", putTodo)
	router.DELETE("api/todo", deleteTodo)

	// Swagger configuration (that will overwrites the General API annotations on main().
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "TODO API"
	docs.SwaggerInfo.Description = "TODO API sample by Gin"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8001"
	router.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))
	// url := swagger.URL("https://petstore.swagger.io/v2/swagger.json") // The url pointing to API definition
	// router.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler, url))

	router.Run("localhost:8001")
}

// @Title Get TODO list
// @Description The handler to response the TODO list
// @Router /api/todo [get]
// @Accept json
// @Produce json
// @Success 200 {object} types.TodoPageData "OK"
// getTodoList: The handler to response the TODO list
func getTodoList(c *gin.Context) {

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, myTodoList)
}

// @Title Get a TODO by its Id
// @Description The handler for getting the TODO by Id
// @Router /api/todo/{id} [get]
// @Param id path string true "A TODO's Id."
// @Accept json
// @Produce json
// @Success 200 {object} types.Todo "OK"
// @Success 204 "No Content"
// getTodo: The handler for getting the TODO by Id
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

// @Title Search TODOs
// @Description The handler for searching the TODOs by Title and IsDone
// @Router /api/todo/search [get]
// @Param title query string false "Contained keyword for TODO's Title."
// @Param isDone query boolean false "Matched value for TODO's IsDone." default(false)
// @Accept json
// @Produce json
// @Success 200 {object} types.TodoPageData "OK"
// searchTodo: The handler for searching the TODOs by title and isDone
func searchTodo(c *gin.Context) {
	queryValIsDone, _ := strconv.ParseBool(c.DefaultQuery("isDone", "false"))
	queryValTitle := c.Query("title")

	matchedTodoList := new(types.TodoPageData)
	matchedTodoList.PageTitle = myTodoList.PageTitle

	for _, todo := range myTodoList.Todos {
		if todo.IsDone == queryValIsDone && strings.Contains(todo.Title, queryValTitle) {
			matchedTodoList.Todos = append(matchedTodoList.Todos, todo)
		}
	}

	// Serialize myTodoList to json and add it to reponse
	c.IndentedJSON(http.StatusOK, matchedTodoList)
}

// @Title Create a new TODO
// @Description The handler to add a new TODO
// @Router /api/todo [post]
// @Param todo body types.TodoPageData true "The new TODO to be created."
// @Accept json
// @Produce json
// @Success 201 {object} types.Todo
// @Failure 400 "Bad Request"
// postTodo: The handler to add a new TODO
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
// @Description The handler to edit a TODO
// @Router /api/todo [put]
// @Param todo body types.Todo true "The TODO to be edited."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
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
// @Description The handler to delete an exist TODO from TODO list
// @Router /api/todo [delete]
// @Param todo body types.Todo true "The TODO to be deleted."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 422 "Unprocessable Entity"
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
