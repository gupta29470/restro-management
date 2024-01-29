package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/middleware"
	"github.com/gupta29470/restro-mngmt/routes"
	"github.com/joho/godotenv"
)

func main() {
	envLoadError := godotenv.Load()
	if envLoadError != nil {
		log.Fatal(envLoadError)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)
}
