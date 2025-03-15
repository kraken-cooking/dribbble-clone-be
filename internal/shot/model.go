package shot

type Shot struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateShotRequest struct {
	Title       string   `json:"title" binding:"required"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}
