package profile

import "time"

type Profile struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateProfileRequest struct {
	Bio     string `json:"bio"`
	Website string `json:"website"`
}
