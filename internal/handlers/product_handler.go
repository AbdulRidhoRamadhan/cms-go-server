package handlers

import (
	"fmt"
	"net/http"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/cloudinary"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

type ProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       *int     `json:"stock"`
	ImgUrl      string   `json:"imgUrl"`
	CategoryID  uint     `json:"categoryId"`
}

const MinimumPrice = 100000 

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(fmt.Errorf("BAD_REQUEST"))
		return
	}

	if request.Name == "" {
		c.Error(fmt.Errorf("NAME_REQUIRED"))
		return
	}

	if request.Description == "" {
		c.Error(fmt.Errorf("DESCRIPTION_REQUIRED"))
		return
	}

	if request.Price == 0 {
		c.Error(fmt.Errorf("PRICE_REQUIRED"))
		return
	}

	if request.Price < MinimumPrice {
		c.Error(fmt.Errorf("PRICE_MIN"))
		return
	}

	if request.CategoryID == 0 {
		c.Error(fmt.Errorf("CATEGORY_REQUIRED"))
		return
	}

	var category models.Category
	if err := h.db.First(&category, uint(request.CategoryID)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("CATEGORY_NOT_FOUND"))
		} else {
			c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		}
		return
	}

	var stockValue int
	if request.Stock == nil {
		stockValue = 0
	} else {
		stockValue = *request.Stock
	}

	userID, _ := c.Get("userID")

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       stockValue,
		CategoryID:  request.CategoryID,
		ImgUrl:      request.ImgUrl,
		AuthorID:    userID.(uint),
	}

	if err := h.db.Create(&product).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	var createdProduct models.Product
	if err := h.db.Preload("Category").
		Preload("Category.Author").
		Preload("Author").
		First(&createdProduct, product.ID).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	var products []models.Product
	
	if err := h.db.Preload("Category").
		Preload("Category.Author").
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, email, role, phone_number, address")
		}).Find(&products).Error; err != nil {
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

	c.JSON(http.StatusOK, formattedProducts)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	
	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	
	var existingProduct models.Product
	if err := h.db.First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("NOT_FOUND"))
			return
		}
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	
	if userRole == "Staff" && existingProduct.AuthorID != userID.(uint) {
		c.Error(fmt.Errorf("FORBIDDEN"))
		return
	}
	
	tx := h.db.Begin()
	
	result := tx.Unscoped().Delete(&models.Product{}, id)
	if result.Error != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		c.Error(fmt.Errorf("NOT_FOUND"))
		return
	}

	var count int64
	if err := tx.Unscoped().Model(&models.Product{}).Count(&count).Error; err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	if count == 0 {
		if err := tx.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1").Error; err != nil {
			tx.Rollback()
			c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product success to delete",
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	var existingProduct models.Product
	if err := h.db.First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("NOT_FOUND"))
			return
		}
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	
	if userRole == "Staff" && existingProduct.AuthorID != userID.(uint) {
		c.Error(fmt.Errorf("FORBIDDEN"))
		return
	}

	var request ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(fmt.Errorf("BAD_REQUEST"))
		return
	}

	updates := make(map[string]interface{})

	if request.Name == "" {
		c.Error(fmt.Errorf("NAME_REQUIRED"))
		return
	}
	updates["name"] = request.Name

	if request.Description == "" {
		c.Error(fmt.Errorf("DESCRIPTION_REQUIRED"))
		return
	}
	updates["description"] = request.Description

	if request.ImgUrl != "" {
		updates["img_url"] = request.ImgUrl
	}

	if request.Price == 0 {
		c.Error(fmt.Errorf("PRICE_REQUIRED"))
		return
	}
	if request.Price < MinimumPrice {
		c.Error(fmt.Errorf("PRICE_MIN"))
		return
	}
	updates["price"] = request.Price

	if request.Stock == nil {
		updates["stock"] = 0
	} else {
		updates["stock"] = *request.Stock
	}

	if request.CategoryID == 0 {
		c.Error(fmt.Errorf("CATEGORY_REQUIRED"))
		return
	}
	
	var category models.Category
	if err := h.db.First(&category, request.CategoryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("CATEGORY_NOT_FOUND"))
		} else {
			c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		}
		return
	}
	updates["category_id"] = request.CategoryID

	result := h.db.Model(&models.Product{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	if result.RowsAffected == 0 {
		c.Error(fmt.Errorf("NOT_FOUND"))
		return
	}

	var updatedProduct models.Product
	if err := h.db.Preload("Category").
		Preload("Category.Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, email, role, phone_number, address, created_at, updated_at")
		}).
		Preload("Author").
		First(&updatedProduct, id).Error; err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) UploadImage(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := h.db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(fmt.Errorf("NOT_FOUND"))
			return
		}
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	if userRole == "Staff" && product.AuthorID != userID.(uint) {
		c.Error(fmt.Errorf("FORBIDDEN"))
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.Error(fmt.Errorf("IMAGE_REQUIRED"))
		return
	}

	if !isValidImageType(file.Header.Get("Content-Type")) {
		c.Error(fmt.Errorf("INVALID_IMAGE_TYPE"))
		return
	}

	imageURL, err := cloudinary.UploadImage(file)
	if err != nil {
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	tx := h.db.Begin()
	
	if err := tx.Model(&product).Update("img_url", imageURL).Error; err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("INTERNAL_SERVER_ERROR"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image Product success to update",
		"imgUrl": imageURL,
	})
}

func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}
	return validTypes[contentType]
}
