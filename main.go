/**
 * @author Chi-Chen Chang (ccchang915@gmail.com)
 */

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simple-todo-list/controllers"
	"simple-todo-list/db"
	_ "simple-todo-list/docs"
)

//	@title			TODO APIs
//	@version		1.0
//	@description	API documentation for the todo list service.

//	@contact.name	Chi-Chen Chang
//	@contact.email	ccchang915@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@schemes		http
func main() {
	app := gin.Default()
	if err := db.InitializeMySQL(); err != nil {
		panic(err)
	}

	initRoutes(app)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func initRoutes(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO APIs
	app.GET("/todos", controllers.GetTasks)
	app.GET("/todos/:id", controllers.GetTask)
	app.POST("/todo", controllers.CreateTask)
	app.PUT("/todos/:id", controllers.UpdateTask)
	app.DELETE("/todos/:id", controllers.DeleteTask)
}
