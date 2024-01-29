package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}
