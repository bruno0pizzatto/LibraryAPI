package controllers

import (
	"strconv"
	"github.com/bruno0pizzatto/LibraryAPI/database"
	"github.com/bruno0pizzatto/LibraryAPI/models"
	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"Error": "ID has to be integer"})
		return
	}
	db := database.GetDataBase()
	var user models.User
	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": "Cannot find user: " + err.Error()})
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot bind JSON " + err.Error()})
		return
	}
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": "Cannot create user: " + err.Error()})
		return
	}
	c.JSON(200, user)
}

func ShowUsers(c *gin.Context) {
	db := database.GetDataBase()
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": "Cannot list users: " + err.Error()})
		return
	}
	c.JSON(200, users)
}
