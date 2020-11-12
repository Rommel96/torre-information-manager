package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rommel96/torre-information-manager/backend/models"
)

const bearer string = "Bearer "

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
				Status:  models.StatusError,
				Message: models.AuthHeaderNotFound,
			})
			return
		}
		tokenString := authHeader[len(bearer):]
		jwtToken, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.MsgResponse{
				Status:  models.StatusError,
				Message: err.Error(),
			})
			return
		}
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.MsgResponse{
				Status:  models.StatusError,
				Message: models.InvalidToken,
			})
			return
		} else {
			c.Set("email", claims["email"])
		}
		c.Next()
	}
}
