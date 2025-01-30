package middleware

import (
	"fmt"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorizationForAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		if user.Role != "Admin" {
			c.Error(fmt.Errorf("FORBIDDEN"))
			c.Abort()
			return
		}

		c.Next()
	}
} 