package route

import (
	"section22_buku/constants"
	"section22_buku/controller"

	m "section22/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// buku
	e.POST("/buku/login", controller.LoginUsersController)

	//jwt
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/buku/jwt", controller.GetBuku)
	eJwt.POST("/buku/jwt", controller.CreateBuku)
	eJwt.GET("/buku/jwt/:id", controller.GetBukuByIdAuth)
	eJwt.DELETE("/buku/jwt/:id", controller.DeleteBuku)
	eJwt.PUT("/buku/jwt/:id", controller.UpdateBuku)

	return e
}
