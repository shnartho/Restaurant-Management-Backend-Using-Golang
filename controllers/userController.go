package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func SingUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
