package field

import (
	"gorm.io/gorm"
)

const table = "field"

type FieldRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *FieldRepo {
	return &FieldRepo{
		db: db,
	}
}

func (ur *FieldRepo) Save(field Field) bool {
	result := ur.db.Table(table).Save(&field)

	return result.RowsAffected > 0
}
