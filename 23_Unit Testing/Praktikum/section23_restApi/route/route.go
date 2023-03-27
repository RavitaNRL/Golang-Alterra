package route

import (
	"section23_restAPI/controller"

	"github.com/labstack/echo"

	"section23_restAPI/constants"
	m "section23_restAPI/middleware"

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
	eJwt.Use(mid.JWT([]byte(constants.SECRETE_JWT)))
	eJwt.GET("/user/jwt", controller.GetUsers)
	eJwt.GET("/user/jwt/:id", controller.GetUserById)
	eJwt.DELETE("/user/jwt/:id", controller.DeleteUser)
	eJwt.PUT("/user/jwt/:id", controller.UpdateUser)

	return e

}
