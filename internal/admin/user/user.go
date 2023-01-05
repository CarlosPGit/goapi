package user

type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Country  string `json:"country"`
}
