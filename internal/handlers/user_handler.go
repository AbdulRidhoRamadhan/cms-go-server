package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/hash"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) Register(c *gin.Context) {
	var registerRequest struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		Role        string `json:"role"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.Error(fmt.Errorf("BAD_REQUEST"))
		return
	}

	if registerRequest.Username == "" {
		c.Error(fmt.Errorf("USERNAME_REQUIRED"))
		return
	}

	if registerRequest.Email == "" {
		c.Error(fmt.Errorf("EMAIL_REQUIRED"))
		return
	}

	if !regexp.MustCompile(emailRegex).MatchString(registerRequest.Email) {
		c.Error(fmt.Errorf("INVALID_EMAIL_FORMAT"))
		return
	}

	if registerRequest.Password == "" {
		c.Error(fmt.Errorf("PASS_REQUIRED"))
		return
	}

	if len(registerRequest.Password) < 5 {
		c.Error(fmt.Errorf("PASSWORD_LENGTH"))
		return
	}

	var existingUser models.User
	if err := h.db.Where("email = ?", registerRequest.Email).First(&existingUser).Error; err == nil {
		c.Error(fmt.Errorf("EMAIL_ALREADY_EXISTS"))
		return
	}

	role := "Staff"
	if strings.ToLower(registerRequest.Role) == "admin" {
		role = "Admin"
	}

	user := models.User{
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		Password:    registerRequest.Password,
		PhoneNumber: registerRequest.PhoneNumber,
		Address:     registerRequest.Address,
		Role:        role,
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"email":       user.Email,
		"phoneNumber": user.PhoneNumber,
		"address":     user.Address,
		"role":        user.Role,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.Error(fmt.Errorf("UNAUTHENTICATED"))
		return
	}

	if loginData.Email == "" {
		c.Error(fmt.Errorf("EMAIL_REQUIRED"))
		return
	}

	if !regexp.MustCompile(emailRegex).MatchString(loginData.Email) {
		c.Error(fmt.Errorf("INVALID_EMAIL_FORMAT"))
		return
	}

	if loginData.Password == "" {
		c.Error(fmt.Errorf("PASS_REQUIRED"))
		return
	}

	var user models.User
	if err := h.db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		log.Printf("User not found: %v", err)
		c.Error(fmt.Errorf("UNAUTHENTICATED"))
		return
	}

	log.Printf("Found user: %+v", user)
	log.Printf("Comparing passwords - Input: %s, Stored: %s", loginData.Password, user.Password)
	
	if !hash.CheckPassword(loginData.Password, user.Password) {
		log.Printf("Password mismatch")
		c.Error(fmt.Errorf("UNAUTHENTICATED"))
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"role":        user.Role,
	})
}
