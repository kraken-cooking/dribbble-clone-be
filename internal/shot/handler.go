package shot

import (
	"dribbble-clone-be/internal/middleware"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	// Clean and validate tags
	var cleanTags []string
	for _, tag := range req.Tags {
		if trimmed := strings.TrimSpace(tag); trimmed != "" {
			cleanTags = append(cleanTags, trimmed)
		}
	}

	// Create the shot
	shot := Shot{
		UserID:      userID,
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		ImageURL:    req.ImageURL,
		Tags:        strings.Join(cleanTags, ","),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if result := h.db.Create(&shot); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shot"})
		return
	}

	// Prepare response
	response := ShotResponse{
		ID:          shot.ID,
		UserID:      shot.UserID,
		Title:       shot.Title,
		ImageURL:    shot.ImageURL,
		Description: shot.Description,
		Tags:        cleanTags,
		CreatedAt:   shot.CreatedAt,
		UpdatedAt:   shot.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Shot created successfully",
		"data":    response,
	})
}

type PaginationResponse struct {
	Total       int64          `json:"total"`
	CurrentPage int            `json:"current_page"`
	PageSize    int            `json:"page_size"`
	TotalPages  int            `json:"total_pages"`
	Data        []ShotResponse `json:"data"`
}

func (h *Handler) GetShots(c *gin.Context) {
	// Get query parameters
	tag := c.Query("tag")
	page := 1
	pageSize := 10

	// Parse pagination parameters
	if pageStr := c.Query("page"); pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}
	if sizeStr := c.Query("page_size"); sizeStr != "" {
		if parsedSize, err := strconv.Atoi(sizeStr); err == nil && parsedSize > 0 && parsedSize <= 50 {
			pageSize = parsedSize
		}
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Initialize the query
	query := h.db.Model(&Shot{})

	// Apply tag filter if provided
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count shots"})
		return
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// Get paginated shots
	var shots []Shot
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&shots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shots"})
		return
	}

	// Convert shots to response format
	var responses []ShotResponse
	for _, shot := range shots {
		var tags []string
		if shot.Tags != "" {
			tags = strings.Split(shot.Tags, ",")
		}

		responses = append(responses, ShotResponse{
			ID:          shot.ID,
			UserID:      shot.UserID,
			Title:       shot.Title,
			ImageURL:    shot.ImageURL,
			Description: shot.Description,
			Tags:        tags,
			CreatedAt:   shot.CreatedAt,
			UpdatedAt:   shot.UpdatedAt,
		})
	}

	// Prepare pagination response
	response := PaginationResponse{
		Total:       total,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		Data:        responses,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shots retrieved successfully",
		"data":    response,
	})
}

func (h *Handler) GetShot(c *gin.Context) {
	shotID := c.Param("id")

	var shot Shot
	if result := h.db.First(&shot, shotID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shot not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shot"})
		return
	}

	// Parse tags from comma-separated string
	var tags []string
	if shot.Tags != "" {
		tags = strings.Split(shot.Tags, ",")
	}

	response := ShotResponse{
		ID:          shot.ID,
		UserID:      shot.UserID,
		Title:       shot.Title,
		ImageURL:    shot.ImageURL,
		Description: shot.Description,
		Tags:        tags,
		CreatedAt:   shot.CreatedAt,
		UpdatedAt:   shot.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shot retrieved successfully",
		"shot":    response,
	})
}
