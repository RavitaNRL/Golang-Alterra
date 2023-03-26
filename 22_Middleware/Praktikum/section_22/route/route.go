package route

import (
	"section_22/constants"
	"section_22/controller"

	"github.com/labstack/echo"

	m "section_22/middleware"

	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// // buku
	// e.GET("/buku", controller.GetBuku)
	// e.GET("/buku/:id", controller.GetBukuById)
	// e.POST("/buku", controller.CreateBuku)
	// e.DELETE("/buku/:id", controller.DeleteBuku)
	// e.PUT("/buku/:id", controller.UpdateBuku)

	// user
	e.GET("/user", controller.GetUsers)
	e.GET("/user/:id", controller.GetUserById)
	e.POST("/user", controller.CreateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.POST("/user/login", controller.LoginUser)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/user", controller.GetUsers)
	// eJwt.GET("/user/jwt/:id", controller.GetUserByToken)
	eJwt.GET("/user/:id", controller.GetUserByToken)
	eJwt.DELETE("/user/:id", controller.DeleteUserByToken)
	eJwt.PUT("/user/:id", controller.UpdateUserByToken)

	return e

}
