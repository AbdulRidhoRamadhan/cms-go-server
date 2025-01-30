package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublicHandler struct {
	db *gorm.DB
}

func NewPublicHandler(db *gorm.DB) *PublicHandler {
	return &PublicHandler{db: db}
}

func (h *PublicHandler) GetAllProducts(c *gin.Context) {
	var products []models.Product
	
	
	search := c.Query("search")
	sort := c.Query("sort")
	categoryId := c.Query("categoryId")
	categoryName := c.Query("categoryName")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	
	query := h.db

	
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	
	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}

	
	if categoryName != "" {
		query = query.Joins("JOIN categories ON products.category_id = categories.id").
			Where("categories.name ILIKE ?", "%"+categoryName+"%")
	}

	
	if sort == "asc" {
		query = query.Order("created_at ASC")
	} else if sort == "desc" {
		query = query.Order("created_at DESC")
	}

	
	var count int64
	if err := query.Model(&models.Product{}).Count(&count).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	
	if err := query.Preload("Category").
		Preload("Category.Author").
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, email, role, phone_number, address")
		}).
		Limit(limitInt).
		Offset(offset).
		Find(&products).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	
	formattedProducts := make([]gin.H, len(products))
	for i, product := range products {
		formattedProducts[i] = gin.H{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"imgUrl":      product.ImgUrl,
			"categoryId":  product.CategoryID,
			"authorId":    product.AuthorID,
			"Category": gin.H{
				"id":       product.Category.ID,
				"name":     product.Category.Name,
				"authorId": product.Category.AuthorID,
				"Author":   product.Category.Author,
			},
			"User":      product.Author,
			"createdAt": product.CreatedAt,
			"updatedAt": product.UpdatedAt,
		}
	}

	totalPages := int(math.Ceil(float64(count) / float64(limitInt)))

	c.JSON(http.StatusOK, gin.H{
		"products":    formattedProducts,
		"totalPages": totalPages,
		"currentPage": pageInt,
	})
}

func (h *PublicHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.db.Preload("Category").
		Preload("Category.Author").
		Preload("Author").
		First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("NOT_FOUND"))
			return
		}
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	
	formattedProduct := gin.H{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"imgUrl":      product.ImgUrl,
		"categoryId":  product.CategoryID,
		"authorId":    product.AuthorID,
		"Category": gin.H{
			"id":       product.Category.ID,
			"name":     product.Category.Name,
			"authorId": product.Category.AuthorID,
			"Author":   product.Category.Author,
		},
		"User":      product.Author,
		"createdAt": product.CreatedAt,
		"updatedAt": product.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"productById": formattedProduct,
	})
}
