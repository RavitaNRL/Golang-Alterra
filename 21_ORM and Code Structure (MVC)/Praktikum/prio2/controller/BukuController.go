package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"prio2/config"
	"prio2/models"
)

// get all buku
func GetBuku(c echo.Context) error {
	//inisialisasi
	var buku []models.Buku

	//pemanggilan data dari database
	err := config.DB.Find(&buku).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"buku":    buku,
	})
}

// get buku by id
func GetBukuById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	buku := models.Buku{}
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get buku by id",
		"buku":    buku,
	})
}

// create buku
func CreateBuku(c echo.Context) error {
	buku := models.Buku{}
	c.Bind(&buku)

	if err := config.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create buku",
		"buku":    buku,
	})
}

// update buku by id
func UpdateBuku(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	buku := models.Buku{}
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&buku)

	if err := config.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update buku",
		"buku":    buku,
	})
}

// delete buku by id
func DeleteBuku(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	buku := models.Buku{}
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete buku",
	})
}

//get all users
func GetUsers(c echo.Context) error {
	//inisialisasi
	var users []models.User

	//pemanggilan data dari database
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}

//get user by id
func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

//create user
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

//update user by id
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

//delete user by id
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}
