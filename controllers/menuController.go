package controller

import (
	"github.com/gin-gonic/gin"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}
