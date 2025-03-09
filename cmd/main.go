package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"your-module/internal/auth"
	"your-module/internal/profile"
	"your-module/internal/shot"
	"your-module/internal/middleware"
	"your-module/pkg/database"
)

func main() {
	// Initialize DB
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()

	// Auth routes
	authHandler := auth.NewHandler(db)
	router.POST("/auth/signup", authHandler.Signup)
	router.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Profile routes
		profileHandler := profile.NewHandler(db)
		protected.GET("/profile", profileHandler.GetProfile)
		protected.PUT("/profile", profileHandler.UpdateProfile)

		// Shot routes
		shotHandler := shot.NewHandler(db)
		protected.POST("/shots", shotHandler.UploadShot)
		protected.GET("/shots", shotHandler.GetShots)
		protected.GET("/shots/:id", shotHandler.GetShot)
	}

	log.Fatal(router.Run(":8080"))
} 