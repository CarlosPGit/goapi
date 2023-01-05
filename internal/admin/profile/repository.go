package profile

import (
	"gorm.io/gorm"
)

const table = "profile"

type ProfiledRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *ProfiledRepo {
	return &ProfiledRepo{
		db: db,
	}
}

func (ur *ProfiledRepo) Save(profile Profile) bool {
	result := ur.db.Table(table).Save(&profile)

	return result.RowsAffected > 0
}
