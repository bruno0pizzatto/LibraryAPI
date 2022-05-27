package controllers

import (
	"strconv"

	"github.com/bruno0pizzatto/LibraryAPI/database"
	"github.com/bruno0pizzatto/LibraryAPI/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID has to be integer",
		})
	}

	db := database.GetDataBase()
	var book models.Book

	err = db.First(book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON " + err.Error(),
		})

		return
	}
	err = db.Create(book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {

}
