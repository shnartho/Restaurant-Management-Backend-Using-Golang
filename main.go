package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
