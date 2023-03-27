package controller

import (
	"net/http"
	"section22_buku/config"
	"section22_buku/middleware"
	"section22_buku/models"
	"strconv"

	"github.com/labstack/echo"
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

func LoginUsersController(c echo.Context) error {
	buku := models.Buku{}
	c.Bind(&buku)

	if err := config.DB.Where("email = ? AND password = ?", buku.Email, buku.Password).First(&buku).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login user",
			"status":  err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(buku.ID), buku.Judul)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	buku.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   buku,
	})
}

// Get buku by ID using JWT authentication
func GetBukuByIdAuth(c echo.Context) error {
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

// create buku by JWT Authentication
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

// update buku by id using JWT Authentication
func UpdateBuku(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Book ID")
	}

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

// delete buku by id using JWT Authentication
func DeleteBuku(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Book ID")
	}

	buku := models.Buku{}
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete buku",
		"buku":    buku,
	})
}
