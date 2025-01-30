package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			var status int
			var message string

			switch err.Error() {
			case "BAD_REQUEST":
				status = http.StatusBadRequest
				message = "Invalid input format"
			case "NOT_FOUND":
				status = http.StatusNotFound
				message = "Data not found"

			// User
			case "USERNAME_REQUIRED":
				status = http.StatusBadRequest
				message = "Username is required"
			case "PASSWORD_LENGTH":
				status = http.StatusBadRequest
				message = "Password must be at least 5 characters"
			case "EMAIL_REQUIRED":
				status = http.StatusBadRequest
				message = "Email is required"
			case "EMAIL_ALREADY_EXISTS":
				status = http.StatusBadRequest
				message = "Email already exists"
			case "PASS_REQUIRED":
				status = http.StatusBadRequest
				message = "Password is required"
			case "UNAUTHENTICATED":
				status = http.StatusUnauthorized
				message = "Invalid email or password"
			case "UNAUTHORIZED":
				status = http.StatusUnauthorized
				message = "Invalid token"
			case "FORBIDDEN":
				status = http.StatusForbidden
				message = "You don't have access"

			// Product
			case "PRICE_MIN":
				status = http.StatusBadRequest
				message = "Minimum price is Rp.100.000,00"
			case "CATEGORY_NOT_FOUND":
				status = http.StatusBadRequest
				message = "Category not found"
			case "NAME_REQUIRED":
				status = http.StatusBadRequest
				message = "Product name is required"
			case "DESCRIPTION_REQUIRED":
				status = http.StatusBadRequest
				message = "Product description is required"
			case "PRICE_REQUIRED":
				status = http.StatusBadRequest
				message = "Product price is required"
			case "CATEGORY_REQUIRED":
				status = http.StatusBadRequest
				message = "Category is required"
			case "IMAGE_REQUIRED":
				status = http.StatusBadRequest
				message = "Product image is required"
			case "INVALID_IMAGE_TYPE":
				status = http.StatusBadRequest
				message = "Invalid image type. Supported types: JPEG, PNG, GIF, WEBP"
			default:
				status = http.StatusInternalServerError
				message = "Internal Server Error"
			}

			c.AbortWithStatusJSON(status, gin.H{"message": message})
		}
	}
}
