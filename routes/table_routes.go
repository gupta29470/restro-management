package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
