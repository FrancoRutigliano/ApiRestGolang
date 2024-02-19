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
	app.GET("/books", controllers.ObtenerLibros)
	app.GET("/books/:id", controllers.ObetenerLibroById)
	app.POST("/books", controllers.CrearLibro)
	app.PUT("/books/:id", controllers.ModificarLibro)
	app.DELETE("/books/:id", controllers.EliminarLibro)

	return app
}
