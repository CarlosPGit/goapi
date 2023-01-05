package profile

type Profile struct {
	UserID   int    `gorm:"primaryKey"`
	Position int    `json:"position"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
}
