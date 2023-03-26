package route

import (
	"section_22_book/constants"
	"section_22_book/controller"

	"github.com/labstack/echo"

	m "section_22_book/middleware"

	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// buku
	e.GET("/buku", controller.GetBuku)
	e.GET("/buku/:id", controller.GetBukuById)
	e.POST("/buku", controller.CreateBuku)
	e.DELETE("/buku/:id", controller.DeleteBuku)
	e.PUT("/buku/:id", controller.UpdateBuku)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/buku", controller.GetBuku)
	eJwt.GET("/buku/:id", controller.GetBukuByIdAuth)
	eJwt.POST("/buku", controller.CreateBukuAuth)
	eJwt.DELETE("/buku/:id", controller.DeleteBukubyIdAuth)
	eJwt.PUT("/buku/:id", controller.UpdateBukubyIdAuth)

	return e
}
