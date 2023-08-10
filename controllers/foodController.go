package controller

import (
	"context"
	"time"
	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second))
		foodId := c.Param("food_id")
		var food models.Food
		foodCollection.FoodOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}
