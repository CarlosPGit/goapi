package user

import (
	"gorm.io/gorm"
)

const table = "user"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) Save(user User) bool {
	result := ur.db.Table(table).Save(&user)

	return result.RowsAffected > 0
}

func (ur *UserRepo) FindByEmail(email string) User {
	var user User
	ur.db.Table(table).Find(&user, "email = ?", email)

	return user
}
