package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
