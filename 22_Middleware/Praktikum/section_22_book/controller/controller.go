package controller

import (
	"net/http"
	"section_22_book/config"
	"section_22_book/models"
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

// get buku by id authentication jwt
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

//crete buku authentication jwt
func CreateBukuAuth(c echo.Context) error {
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

// update buku by id authentication jwt
func UpdateBukubyIdAuth(c echo.Context) error {
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

// delete buku by id authentication jwt
func DeleteBukubyIdAuth(c echo.Context) error {
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
