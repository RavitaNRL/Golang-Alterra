package main

import (
	"section24/config"
	"section24/constants"
	"section24/controller"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func main() {

	// untuk connect ke database menggunakan fungsi ConnectDB dari config
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	// membuat aplikasi echo
	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))
	app.POST("users/login", controller.LoginUsersController(db))

	// untuk menjalankan authentikasi jwt
	appJwt := app.Group("")
	appJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	appJwt.GET("/users/jwt", controller.GetAllUsers(db))

	app.Start(":8080")
}
