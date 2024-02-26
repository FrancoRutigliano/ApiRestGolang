package routes

import (
	"github.com/FrancoRutigliano/controllers"
	"github.com/gin-gonic/gin"
)

func SetUp(app *gin.Engine) {

	// ROUTES
	app.GET("/books", controllers.GetBooks)
	app.GET("/books/:id", controllers.GetBookById)
	app.POST("/books", controllers.CreateBook)
	app.PUT("/books/:id", controllers.ModifiedBook)
	app.DELETE("/books/:id", controllers.DeleteBook)

}
