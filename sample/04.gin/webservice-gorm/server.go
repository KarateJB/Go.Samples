package main

import (
	"example/webservice/docs"
	dbservice "example/webservice/services/db"
	todoservice "example/webservice/services/todo"
	userservice "example/webservice/services/user"
	types "example/webservice/types/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm/logger"
)

const HTTP_HEADER_ROW_COUNT = "X-Row-Count"

var userService *userservice.UserAccess
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
			userRg.GET(":id", getUser)
			userRg.POST("", postUser)
			userRg.PUT("", putUser)
			userRg.DELETE("", deleteUser)
		}
		todoRg := apiRouterGroup.Group("/todo")
		{
			todoRg.GET(":id", getTodo) // The id is required for matching this routing
			// todoRg.GET("*id", getTodoById) // The id is optional for matching this routing, e.q. api/todo/ or api/todo/xxx
			todoRg.GET("search", searchTodo)
			todoRg.POST("", postTodo)
			todoRg.PUT("", putTodo)
			todoRg.DELETE("", deleteTodo)
		}
		todosRg := apiRouterGroup.Group("/todos")
		{
			todosRg.GET("", getAllTodos)
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
	dsn := "host=localhost user=postgres password=1qaz2wsx dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	dbService := dbservice.New(dsn, logger.Info)
	dbService.Migrate()
	dbService.InitData()

	// Init services
	userService = userservice.New(dbService.DB)
	todoService = todoservice.New(dbService.DB)

	router.Run("localhost:8001")
}

// @Tags User
// @Title Get the User by Id
// @Description The handler for getting the User by Id
// @Router /api/user/{id} [get]
// @Param id path string true "The User's Id."
// @Accept json
// @Produce json
// @Success 200 {object} types.User "OK"
// @Success 204 "No Content"
func getUser(c *gin.Context) {
	id := c.Param("id") // Get the value from api/user/:id

	if user := userService.Get(id); user == nil {
		c.Writer.WriteHeader(http.StatusNoContent) // If not found, response 204
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
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

// @Tags Todo
// @Title Search TODOs
// @Description The handler for searching the TODOs by Title and IsDone
// @Router /api/todo/search [get]
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

// @Tags User
// @Title Create a new User
// @Description The handler to add a new User
// @Router /api/user [post]
// @Param user body types.User true "The new User to be created."
// @Accept json
// @Produce json
// @Success 201 {object} types.User
// @Failure 400 "Bad Request"
func postUser(c *gin.Context) {
	var newUser types.User
	if err := c.BindJSON(&newUser); err != nil {
		// return
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	userService.Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
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

// @Tags User
// @Title Edit a User
// @Description The handler to edit a User
// @Router /api/user [put]
// @Param user body types.Todo true "The User to be edited."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 422 "Unprocessable Entity"
func putUser(c *gin.Context) {
	var editUser types.User
	if err := c.BindJSON(&editUser); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := userService.Update(&editUser); count == 0 {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	}
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

// @Tags User
// @Title Delete a User
// @Description The handler to delete an exist User from User list
// @Router /api/user [delete]
// @Param todo body types.Todo true "The User to be deleted."
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 422 "Unprocessable Entity"
func deleteUser(c *gin.Context) {
	var deleteUser types.User
	if err := c.BindJSON(&deleteUser); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := userService.Delete(&deleteUser); count == 0 {
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
