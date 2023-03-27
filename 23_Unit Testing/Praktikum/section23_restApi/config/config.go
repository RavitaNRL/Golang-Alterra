package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"section23_restAPI/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/section23_unittesting?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigration()
}

func initMigration() {
	DB.AutoMigrate(&models.User{})
}

func InitDBTest() {
	dsn := "root:@tcp(localhost:3306)/section23_testing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigreTest()
}

func InitMigreTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
}
