package main

import (
	todoapi "example/webservice/api/todo"
	userapi "example/webservice/api/user"
	"example/webservice/config"
	"example/webservice/docs"
	dbservice "example/webservice/services/db"
	todoservice "example/webservice/services/todo"
	"fmt"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

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
			todoRg.GET(":id", todoapi.GetTodo) // The id is required for matching this routing
			// todoRg.GET("*id", todoapi.GetTodoById) // The id is optional for matching this routing, e.q. api/todo/ or api/todo/xxx
			todoRg.POST("", todoapi.PostTodo)
			todoRg.PUT("", todoapi.PutTodo)
			todoRg.DELETE("", todoapi.DeleteTodo)
		}
		todosRg := apiRouterGroup.Group("/todos")
		{
			todosRg.GET("", todoapi.GetAllTodos)
			todosRg.GET("search", todoapi.SearchTodo)
			todosRg.DELETE("", todoapi.DeleteTodos)
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

	// Start listening
	configs := config.Init()
	router.Run(fmt.Sprintf("localhost:%s", configs.Port))
}
