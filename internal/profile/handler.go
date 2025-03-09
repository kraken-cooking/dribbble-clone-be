package profile

import (
	"dribbble-clone-be/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetProfile(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// TODO: Implement get profile logic
	log.Printf("Get profile successful, useId: %d", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Get profile successful"})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// TODO: Implement update profile logic
	log.Printf("Get profile successful, useId: %d", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Update profile successful"})
}
