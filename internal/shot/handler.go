package shot

import (
	"dribbble-clone-be/internal/middleware"
	"log"
	"net/http"
	"strings"

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

	var req CreateShotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	shot := Shot{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Tags:        strings.Join(req.Tags, ","),
	}

	if result := h.db.Create(&shot); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shot created successfully", "shot": shot})
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
