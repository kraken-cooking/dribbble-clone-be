package shot

import "time"

type Shot struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateShotRequest struct {
	Title       string   `json:"title" binding:"required"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url" binding:"required"`
}

type ShotResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
