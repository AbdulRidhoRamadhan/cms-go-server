package middleware

import (
	"fmt"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, _ := c.Get("userID")
		path := c.FullPath()

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.Error(fmt.Errorf("UNAUTHORIZED"))
			c.Abort()
			return
		}

		if user.Role == "Admin" {
			c.Next()
			return
		}

		if path == "/categories/:id" {
			var category models.Category
			if err := db.First(&category, id).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					c.Error(fmt.Errorf("NOT_FOUND"))
					c.Abort()
					return
				}
				c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
				c.Abort()
				return
			}

			if category.AuthorID != user.ID {
				c.Error(fmt.Errorf("FORBIDDEN"))
				c.Abort()
				return
			}
		} else {
			var product models.Product
			if err := db.First(&product, id).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					c.Error(fmt.Errorf("NOT_FOUND"))
					c.Abort()
					return
				}
				c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
				c.Abort()
				return
			}

			if product.AuthorID != user.ID {
				c.Error(fmt.Errorf("FORBIDDEN"))
				c.Abort()
				return
			}
		}

		c.Next()
	}
} 