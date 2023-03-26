package controller

import (
	"net/http"
	"section_22/config"
	"section_22/constants"
	"section_22/middleware"
	"section_22/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// // get all buku
// func GetBuku(c echo.Context) error {
// 	//inisialisasi
// 	var buku []models.Buku

// 	//pemanggilan data dari database
// 	err := config.DB.Find(&buku).Error
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success",
// 		"buku":    buku,
// 	})
// }

// // get buku by id
// func GetBukuById(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	buku := models.Buku{}
// 	if err := config.DB.First(&buku, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get buku by id",
// 		"buku":    buku,
// 	})
// }

// // create buku
// func CreateBuku(c echo.Context) error {
// 	buku := models.Buku{}
// 	c.Bind(&buku)

// 	if err := config.DB.Save(&buku).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create buku",
// 		"buku":    buku,
// 	})
// }

// // update buku by id
// func UpdateBuku(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	buku := models.Buku{}
// 	if err := config.DB.First(&buku, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	c.Bind(&buku)

// 	if err := config.DB.Save(&buku).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success update buku",
// 		"buku":    buku,
// 	})
// }

// // delete buku by id
// func DeleteBuku(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	buku := models.Buku{}
// 	if err := config.DB.First(&buku, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := config.DB.Delete(&buku).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete buku",
// 	})
// }

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

// login
func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"user":    userResponse,
	})
}

// get user by token
// Get user by ID using JWT authentication
func GetUserByToken(c echo.Context) error {
	// Get the user ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	// Get the JWT token from the request header
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
	}

	// Verify the JWT token
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil // Replace with your actual secret key
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}
	if !jwtToken.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	// Get the user data from the database
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Return the user data as a JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// delete user by token
// Delete user by ID using JWT authentication
func DeleteUserByToken(c echo.Context) error {
	// Get the user ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	// Get the JWT token from the request header
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
	}

	// Verify the JWT token
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil // Replace with your actual secret key
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}
	if !jwtToken.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	// Get the user data from the database
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user ID matches the ID of the authenticated user
	userID := int(claims["user_id"].(float64))
	if user.ID != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized to access this resource")
	}

	// Delete the user from the database
	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	// Return a success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

// update user by token
// Update user by ID using JWT authentication
func UpdateUserByToken(c echo.Context) error {
	// Get the user ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	// Get the JWT token from the request header
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
	}

	// Verify the JWT token
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil // Replace with your actual secret key
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}
	if !jwtToken.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	// Get the user data from the database
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user ID matches the ID of the authenticated user
	userID := int(claims["user_id"].(float64))
	if user.ID != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized to access this resource")
	}

	// Parse the JSON request body into a User struct
	requestUser := new(models.User)
	if err := c.Bind(requestUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse request body")
	}

	// Update the user data in the database
	user.Name = requestUser.Name
	user.Email = requestUser.Email
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	// Return the updated user data as a JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}
