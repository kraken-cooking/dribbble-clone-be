package upload

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"dribbble-clone-be/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) UploadImage(c *gin.Context) {
	// Check authentication
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"error": map[string]interface{}{
				"message": "Unauthorized",
				"details": "User not authenticated",
			},
		})
		return
	}

	// Get the file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error": map[string]interface{}{
				"message": "Upload failed",
				"details": "No image file provided",
			},
		})
		return
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	if !allowedExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error": map[string]interface{}{
				"message": "Upload failed",
				"details": "Invalid file type. Allowed types: jpg, jpeg, png, gif",
			},
		})
		return
	}

	// Generate unique filename
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("%d_%d%s", userID, timestamp, ext)
	uploadPath := filepath.Join("uploads", "images", filename)

	// Save the file
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": map[string]interface{}{
				"message": "Upload failed",
				"details": "Failed to save image",
			},
		})
		return
	}

	// Generate file URL
	fileURL := fmt.Sprintf("/uploads/images/%s", filename)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File uploaded successfully",
		"data": gin.H{
			"url":      fileURL,
			"filename": filename,
		},
	})
}
