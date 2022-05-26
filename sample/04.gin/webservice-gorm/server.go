package main

import (
	userapi "example/webservice/api/user"
	"example/webservice/config"
	"example/webservice/docs"
	dbservice "example/webservice/services/db"
	todoservice "example/webservice/services/todo"
	types "example/webservice/types/api"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

const HTTP_HEADER_ROW_COUNT = "X-Row-Count"

var todoService *todoservice.TodoAccess

// @Title TODO API
// @Version 1.0
// @Description TODO API sample by Gin
// @Host localhost:8001
// @BasePath /
// @Contact.Name JB
// @Contact.Url https://karatejb.blogspot.com/
// @Contact.Email xxx@demo.com
func main() {

	/* Init Gin router */
	router := gin.Default()

	// User
	apiRouterGroup := router.Group("/api")
	{
		userRg := apiRouterGroup.Group("/user")
		{
			userRg.GET(":id", userapi.GetUser)
			userRg.POST("", userapi.PostUser)
			userRg.PUT("", userapi.PutUser)
			userRg.DELETE("", userapi.DeleteUser)
		}
		todoRg := apiRouterGroup.Group("/todo")
		{
			todoRg.GET(":id", getTodo) // The id is required for matching this routing
			// todoRg.GET("*id", getTodoById) // The id is optional for matching this routing, e.q. api/todo/ or api/todo/xxx
			todoRg.POST("", postTodo)
			todoRg.PUT("", putTodo)
			todoRg.DELETE("", deleteTodo)
		}
		todosRg := apiRouterGroup.Group("/todos")
		{
			todosRg.GET("", getAllTodos)
			todosRg.GET("search", searchTodo)
			todosRg.DELETE("", deleteTodos)
		}
	}

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

	// DB connect configuration
	dbService := dbservice.New()
	dbService.Migrate()
	dbService.InitData()

	// Init services
	todoService = todoservice.New(dbService.DB)

	configs := config.Init()
	router.Run(fmt.Sprintf("localhost:%s", configs.Port))
}

// @Tags Todos
// @Title Get all TODOs
// @Description The handler to response the TODO list
// @Router /api/todos [get]
// @Accept json
// @Produce json
// @Success 200 {array} types.Todo "OK"
// @Success 204 "No Content"
func getAllTodos(c *gin.Context) {
	if todos := todoService.GetAll(); todos == nil {
		c.Writer.WriteHeader(http.StatusNoContent)
	} else {
		c.IndentedJSON(http.StatusOK, todos)
	}
}

// @Tags Todo
// @Title Get the TODO by its Id
// @Description The handler for getting the TODO by Id
// @Router /api/todo/{id} [get]
// @Param id path string true "A TODO's Id."
// @Accept json
// @Produce json
// @Success 200 {object} types.Todo "OK"
// @Success 204 "No Content"
func getTodo(c *gin.Context) {
	id := c.Param("id") // Get the value from api/todo/:id
	uuid, _ := uuid.Parse(id)

	if todo := todoService.GetOne(uuid); todo == nil {
		c.Writer.WriteHeader(http.StatusNoContent) // If not found, response 204
	} else {
		c.IndentedJSON(http.StatusOK, todo)
	}
}

// @Tags Todos
// @Title Search TODOs
// @Description The handler for searching the TODOs by Title and IsDone
// @Router /api/todos/search [get]
// @Param title query string false "Contained keyword for TODO's Title."
// @Param isDone query boolean false "Matched value for TODO's IsDone." default(false)
// @Accept json
// @Produce json
// @Success 200 {array} types.Todo "OK"
// @Success 204 "No Content"
func searchTodo(c *gin.Context) {
	queryValTitle := c.Query("title")
	queryValIsDone, _ := strconv.ParseBool(c.DefaultQuery("isDone", "false"))

	if todos := todoService.Search(queryValTitle, queryValIsDone); todos == nil {
		c.Writer.WriteHeader(http.StatusNoContent)
	} else {
		c.IndentedJSON(http.StatusOK, todos)
	}
}

// @Tags Todo
// @Title Create a new TODO
// @Description The handler to add a new TODO
// @Router /api/todo [post]
// @Param todo body types.Todo true "The new TODO to be created."
// @Accept json
// @Produce json
// @Success 201 {object} types.Todo
// @Failure 400 "Bad Request"
func postTodo(c *gin.Context) {
	var newTodo types.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		// return
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	entity := todoService.Create(&newTodo)

	// Get the auto-generated Id
	newTodo.Id = entity.Id
	newTodo.TodoExt.Id = entity.TodoExt.Id

	c.IndentedJSON(http.StatusCreated, newTodo)
}

// @Tags Todo
// @Title Edit a TODO
// @Description The handler to edit a TODO
// @Router /api/todo [put]
// @Param todo body types.Todo true "The TODO to be edited."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 422 "Unprocessable Entity"
func putTodo(c *gin.Context) {
	var editTodo types.Todo
	if err := c.BindJSON(&editTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := todoService.Update(&editTodo); count == 0 {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// @Tags Todo
// @Title Delete a TODO
// @Description The handler to delete an TODO
// @Router /api/todo [delete]
// @Param todo body types.Todo true "The TODO to be deleted."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 422 "Unprocessable Entity"
func deleteTodo(c *gin.Context) {
	var deleteTodo types.Todo
	if err := c.BindJSON(&deleteTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := todoService.DeleteOne(&deleteTodo); count == 0 {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// @Tags Todos
// @Title Delete TODOs
// @Description The handler to delete TODOs by their Id
// @Router /api/todos [delete]
// @Param todo body []types.Todo true "The TODOs to be deleted."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
func deleteTodos(c *gin.Context) {
	var deleteTodos []types.Todo
	if err := c.BindJSON(&deleteTodos); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	count := todoService.Delete(&deleteTodos)
	c.Header(HTTP_HEADER_ROW_COUNT, strconv.FormatInt(count, 10))
}
