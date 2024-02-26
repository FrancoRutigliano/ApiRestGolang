package main

import (
	"fmt"
	"log"

	"github.com/FrancoRutigliano/models"
	"github.com/FrancoRutigliano/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := ":8080"
	app := gin.Default()

	routes.SetUp(app)

	models.ConnectDatabase()

	err := app.Run(PORT)
	if err != nil {
		log.Fatal("Error to run the app")
	}

	fmt.Printf("Server running on localhost%s", PORT)
}
