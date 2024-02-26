package controllers

import (
	"net/http"

	"github.com/FrancoRutigliano/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})

}

func CreateBook(c *gin.Context) {

}

func GetBookById(c *gin.Context) {

}

func ModifiedBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
