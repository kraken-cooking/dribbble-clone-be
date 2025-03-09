package profile

type Profile struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatar_url"`
	Website     string `json:"website"`
}

type UpdateProfileRequest struct {
	Bio       string `json:"bio"`
	Website   string `json:"website"`
} 