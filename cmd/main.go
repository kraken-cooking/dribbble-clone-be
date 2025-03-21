package main

import (
	"log"
	"os"
	"time"

	"dribbble-clone-be/internal/auth"
	"dribbble-clone-be/internal/middleware"
	"dribbble-clone-be/internal/profile"
	"dribbble-clone-be/internal/shot"
	"dribbble-clone-be/internal/upload"
	"dribbble-clone-be/pkg/database"

	"github.com/gin-contrib/cors"
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

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve static files
	os.MkdirAll("uploads/images", 0755)
	router.Static("/uploads", "./uploads")

	// Auth routes
	authHandler := auth.NewHandler(db)
	router.POST("/auth/signup", authHandler.Signup)
	router.POST("/auth/login", authHandler.Login)

	// handlers
	shotHandler := shot.NewHandler(db)
	profileHandler := profile.NewHandler(db)
	uploadHandler := upload.NewHandler()

	// Public routes
	router.GET("/shots/:id", shotHandler.GetShot)
	router.GET("/shots", shotHandler.GetShots)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Profile routes
		protected.GET("/profile", profileHandler.GetProfile)
		protected.PUT("/profile", profileHandler.UpdateProfile)

		// Shot routes
		shotHandler := shot.NewHandler(db)
		protected.POST("/shots", shotHandler.UploadShot)

		// Upload routes
		protected.POST("/upload", uploadHandler.UploadImage)
	}

	log.Fatal(router.Run(":8080"))
}
