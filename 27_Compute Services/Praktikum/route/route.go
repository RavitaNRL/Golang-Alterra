package route

import (
	"section22/constants"
	"section22/controller"

	m "section22/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// user
	e.GET("/user", controller.GetUsers)
	e.GET("/user/:id", controller.GetUserById)
	e.POST("/user", controller.CreateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.POST("/user/login", controller.LoginUsersController)

	//jwt
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/user/jwt", controller.GetUsers)
	eJwt.GET("/user/jwt/:id", controller.GetUserById)
	eJwt.DELETE("/user/jwt/:id", controller.DeleteUser)
	eJwt.PUT("/user/jwt/:id", controller.UpdateUser)

	// // buku
	// e.GET("/buku", controller.GetBuku)
	// e.GET("/buku/:id", controller.GetBukuById)
	// e.POST("/buku", controller.CreateBuku)
	// e.DELETE("/buku/:id", controller.DeleteBuku)
	// e.PUT("/buku/:id", controller.UpdateBuku)

	return e

}
