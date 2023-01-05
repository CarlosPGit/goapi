package mysqlclient

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DATABASE_URI string = "capg:7338@tcp(localhost:3306)/sface?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return db, err
}
