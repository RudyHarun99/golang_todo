package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"GO14/finalProject/controllers"
	"GO14/finalProject/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"GO14/finalProject/docs"
)

func Init() {
	db := config.Db
	
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Final Project"
	docs.SwaggerInfo.Description = "Documentation API Todos Final Project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "berhasil",
		})
	})

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.GetTodo)
	r.PUT("/todos/:id", controllers.EditTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	fmt.Println("starting web server at http://localhost:8080/")
	r.Run()
}