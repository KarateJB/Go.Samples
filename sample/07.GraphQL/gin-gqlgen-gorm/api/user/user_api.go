package userapi

import (
	models "example/graphql/graph/model"
	services "example/graphql/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Title Get the User by Id
// @Description The handler for getting the User by Id
// @Router /api/user/{id} [get]
// @Param id path string true "The User's Id."
// @Accept json
// @Produce json
// @Success 200 {object} types.User "OK"
// @Success 204 "No Content"
func GetUser(c *gin.Context) {
	id := c.Param("id") // Get the value from api/user/:id

	if user := services.UserRf.GetOne(id); user == nil {
		c.Writer.WriteHeader(http.StatusNoContent) // If not found, response 204
	} else {
		c.IndentedJSON(http.StatusOK, user)
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
func PostUser(c *gin.Context) {
	var newUser models.NewUser
	if err := c.BindJSON(&newUser); err != nil {
		// return
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	services.UserRf.Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
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
func PutUser(c *gin.Context) {
	var editUser models.EditUser
	if err := c.BindJSON(&editUser); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if _, count := services.UserRf.Update(&editUser); count == 0 {
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
func DeleteUser(c *gin.Context) {
	var deleteUser models.User
	if err := c.BindJSON(&deleteUser); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	if count := services.UserRf.Delete(deleteUser.Id); count == 0 {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	}
}
