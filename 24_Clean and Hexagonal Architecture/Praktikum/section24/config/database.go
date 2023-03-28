package config

import (
	"section24/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// untuk menghubungkan database dengan aplikasi.
func ConnectDB() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// untuk membuat tabel di database.
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
