package shot

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

func (h *Handler) UploadShot(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// TODO: Implement shot upload logic
	log.Printf("Get profile successful, useId: %d", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Shot upload successful"})
}

func (h *Handler) GetShots(c *gin.Context) {
	// TODO: Implement get shots logic
	c.JSON(http.StatusOK, gin.H{"message": "Get shots successful"})
}

func (h *Handler) GetShot(c *gin.Context) {
	shotID := c.Param("id")

	// TODO: Implement get single shot logic
	log.Printf("Get profile successful, useId: %s", shotID)
	c.JSON(http.StatusOK, gin.H{"message": "Get shot successful", "id": shotID})
}
