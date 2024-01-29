package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/controllers"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItems())
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.PATCH("/orderItems/:order_item_id", controllers.UpdateOrderItem())
}
