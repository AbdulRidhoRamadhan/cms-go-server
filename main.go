package main

import (
	"fmt"
	"log"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/config"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/database"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/handlers"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/middleware"
	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/pkg/cloudinary"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	
	cfg := config.LoadConfig()
	
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.InitDB(db)

	if err := cloudinary.InitCloudinary(); err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}

	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.SetTrustedProxies([]string{"127.0.0.1"})

	userHandler := handlers.NewUserHandler(db)
	productHandler := handlers.NewProductHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	publicHandler := handlers.NewPublicHandler(db)

	r.POST("/users/login", userHandler.Login)
	r.GET("/pub", publicHandler.GetAllProducts)
	r.GET("/pub/:id", publicHandler.GetProductByID)

	protected := r.Group("/")
	protected.Use(middleware.Authentication())
	{
		// User routes
		protected.POST("/users/add-user", middleware.AuthorizationForAdmin(db), userHandler.Register)

		// Product routes
		protected.POST("/products", productHandler.CreateProduct)
		protected.GET("/products", productHandler.GetAllProducts)
		protected.GET("/products/:id", productHandler.GetProductByID)
		protected.PUT("/products/:id", middleware.Authorization(db), productHandler.UpdateProduct)
		protected.DELETE("/products/:id", middleware.Authorization(db), productHandler.DeleteProduct)
		protected.PATCH("/products/upload/:id", middleware.Authorization(db), productHandler.UploadImage)

		// Category routes
		protected.POST("/categories", categoryHandler.CreateCategory)
		protected.GET("/categories", categoryHandler.GetAllCategories)
		protected.PUT("/categories/:id", middleware.Authorization(db), categoryHandler.UpdateCategory)
	}

	port := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server running on port %s", port)
	r.Run(port)
}
