package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		clientToken := ginContext.Request.Header.Get("token")
		if clientToken == "" {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"Authentication Middleware(): Token is empty",
				"Token is empty")
			return
		}

		claims := helpers.ValidateTokens(clientToken, ginContext)
		ginContext.Set("email", claims.Email)
		ginContext.Set("first_name", claims.First_Name)
		ginContext.Set("last_name", claims.Last_Name)
		ginContext.Set("phone_number", claims.Phone_Number)
		ginContext.Set("user_id", claims.User_ID)

		ginContext.Next()
	}
}
