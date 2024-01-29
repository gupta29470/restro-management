package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
