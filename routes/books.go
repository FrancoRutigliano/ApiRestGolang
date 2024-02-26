package routes

import (
	"io"
	"os"

	"github.com/FrancoRutigliano/controllers"
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {

	app := gin.New()

	f, _ := os.Create("/log/api.log")

	gin.DefaultWriter = io.MultiWriter(f)

	// ROUTES
	app.GET("/books", controllers.GetBooks)
	app.GET("/books/:id", controllers.GetBookById)
	app.POST("/books", controllers.CreateBook)
	app.PUT("/books/:id", controllers.ModifiedBook)
	app.DELETE("/books/:id", controllers.DeleteBook)

	return app
}
