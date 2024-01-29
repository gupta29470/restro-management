package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/models"
)

func SendCustomErrorHelper(ginContext *gin.Context, statusCode int, message string, actualError string) {
	ginContext.JSON(statusCode, models.CustomError{
		Status_Code: statusCode,
		Message:     message,
		Error:       actualError,
	})
}
