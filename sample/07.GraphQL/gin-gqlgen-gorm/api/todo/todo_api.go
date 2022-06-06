package todoapi

import (
	models "example/graphql/graph/model"
	services "example/graphql/services"
	utils "example/graphql/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Tags Todos
// @Title Get all TODOs
// @Description The handler to response the TODO list
// @Router /api/todos [get]
// @Accept json
// @Produce json
// @Success 200 {array} types.Todo "OK"
// @Success 204 "No Content"
func GetAllTodos(c *gin.Context) {
	if todos := services.TodoRf.GetAll(); todos == nil {
		c.Writer.WriteHeader(http.StatusNoContent)
	} else {
		c.Header(utils.HttpHeaderRowCount, strconv.Itoa(len(todos)))
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
func GetTodo(c *gin.Context) {
	id := c.Param("id") // Get the value from api/todo/:id
	uuid, _ := uuid.Parse(id)

	if todo := services.TodoRf.GetOne(uuid); todo == nil {
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
func SearchTodo(c *gin.Context) {
	queryValTitle := c.Query("title")
	queryValIsDone, _ := strconv.ParseBool(c.DefaultQuery("isDone", "false"))

	if todos := services.TodoRf.Search(queryValTitle, queryValIsDone); todos == nil {
		c.Writer.WriteHeader(http.StatusNoContent)
	} else {
		c.Header(utils.HttpHeaderRowCount, strconv.Itoa(len(*todos)))
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
func PostTodo(c *gin.Context) {
	var newTodo models.NewTodo
	if err := c.BindJSON(&newTodo); err != nil {
		// return
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	createdTodo := services.TodoRf.Create(&newTodo)

	c.IndentedJSON(http.StatusCreated, createdTodo)
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
func PutTodo(c *gin.Context) {
	var editTodo models.EditTodo
	if err := c.BindJSON(&editTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if _, count := services.TodoRf.Update(&editTodo); count == 0 {
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
func DeleteTodo(c *gin.Context) {
	var deleteTodo models.Todo
	if err := c.BindJSON(&deleteTodo); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := services.TodoRf.DeleteOne(deleteTodo.Id); count == 0 {
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
func DeleteTodos(c *gin.Context) {
	var deleteTodos []models.Todo
	if err := c.BindJSON(&deleteTodos); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	todoIds := utils.Map(deleteTodos, func(todo models.Todo) uuid.UUID {
		return todo.Id
	})
	count := services.TodoRf.Delete(&todoIds)
	c.Header(utils.HttpHeaderRowCount, strconv.FormatInt(count, 10))
}
