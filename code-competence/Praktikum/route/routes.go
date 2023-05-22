package route

import (
	"code_competence/constant"
	"code_competence/controller"
	"net/http"

	"github.com/go-playground/validator"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {

	e.Validator = &customValidator{validator: validator.New()}

	// user route
	e.POST("/register", controller.RegisterUserController)
	e.POST("/login", controller.LoginUserController)

	// item route
	items := e.Group("/items", middleware.JWT([]byte(constant.SECRET_JWT)))
	items.POST("", controller.CreateItemController)
	items.GET("", controller.GetAllItemsController)
	items.GET("/id/:id", controller.GetItemByIDController)
	items.PUT("/id/:id", controller.UpdateItemByIDController)
	items.DELETE("/id/:id", controller.DeleteItemByIDController)
	items.GET("/name", controller.GetItemByItemNameController)

	// category route
	category := e.Group("/category", middleware.JWT([]byte(constant.SECRET_JWT)))
	category.POST("", controller.CreateCategoryController)
	category.GET("/:category_id", controller.GetAllByCategoryIDController)
}
