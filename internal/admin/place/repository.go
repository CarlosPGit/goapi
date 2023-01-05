package place

import (
	"gorm.io/gorm"
)

const table = "place"

type PlaceRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *PlaceRepo {
	return &PlaceRepo{
		db: db,
	}
}

func (ur *PlaceRepo) Save(place Place) bool {
	result := ur.db.Table(table).Save(&place)

	return result.RowsAffected > 0
}
