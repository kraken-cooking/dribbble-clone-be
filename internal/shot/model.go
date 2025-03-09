package shot

type Shot struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
}

type CreateShotRequest struct {
	Title    string   `json:"title" binding:"required"`
	Tags     []string `json:"tags"`
} 