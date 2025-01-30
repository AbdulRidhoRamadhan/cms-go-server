package middleware

import (
	"fmt"
	"strings"

	customJWT "github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := customJWT.ValidateToken(tokenString)
		if err != nil {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		c.Set("userID", uint(claims["user_id"].(float64)))
		c.Set("userRole", claims["role"])
		c.Next()
	}
}
