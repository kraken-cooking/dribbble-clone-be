package main

import (
	"log"

	"dribbble-clone-be/internal/auth"
	"dribbble-clone-be/internal/middleware"
	"dribbble-clone-be/internal/profile"
	"dribbble-clone-be/internal/shot"
	"dribbble-clone-be/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
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
