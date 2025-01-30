package handlers

import (
	"fmt"
	"net/http"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.Error(fmt.Errorf("BAD_REQUEST"))
		return
	}

	userID, _ := c.Get("userID")
	category.AuthorID = userID.(uint)

	if err := h.db.Create(&category).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	var createdCategory models.Category
	if err := h.db.Preload("Author").First(&createdCategory, category.ID).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	var categories []models.Category
	
	if err := h.db.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, role, phone_number, address, created_at, updated_at")
	}).Find(&categories).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := h.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("NOT_FOUND"))
			return
		}
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	userID, _ := c.Get("userID")

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.Error(fmt.Errorf("UNAUTHORIZED"))
		return
	}

	if user.Role != "Admin" && category.AuthorID != user.ID {
		c.Error(fmt.Errorf("FORBIDDEN"))
		return
	}

	var updateData models.Category
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.Error(fmt.Errorf("BAD_REQUEST"))
		return
	}

	if err := h.db.Model(&category).Updates(map[string]interface{}{
		"name": updateData.Name,
	}).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}