package controller

import (
	"net/http"
	"prio2/config"
	"section22/middleware"
	"section24/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// untuk memanggil semua data user dengan method GET.
func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := make([]model.User, 0)
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}

}

// untuk membuat data user yang akan terhubung dengan databse menggunakan method POST.
func CreateUser(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		err := db.Create(&user).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}

// untuk login user dengan method POST dan mengembalikan token.
func LoginUsersController(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		c.Bind(&user)

		if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Fail login user",
				"status":  err.Error(),
			})
		}

		token, err := middleware.CreateToken(int(user.ID), user.Name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Fail create JWT TOken",
				"status":  err.Error(),
			})
		}

		user.Token = token

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success login",
			"user":   user,
		})
	}
}
