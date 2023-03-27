package controller

import (
	"net/http"
	"section23_restAPI/constants"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"section23_restAPI/config"
	"section23_restAPI/middleware"
	"section23_restAPI/models"
)

type UserResponse struct {
	Message string
	Data    []models.User
}

// get all users
func GetUsers(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get user by id
func GetUserById(c echo.Context) error {
	// your solution here

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

// delete user by id
func DeleteUser(c echo.Context) error {
	// your solution here

	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success  delete user",
	})

}

// update user by id
func UpdateUser(c echo.Context) error {
	// your solution here

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

func LoginUsersController(c echo.Context) error {
	user := models.User{}
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
		return []byte(constants.SECRETE_JWT), nil // Replace with your actual secret key
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
		return []byte(constants.SECRETE_JWT), nil // Replace with your actual secret key
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
		return []byte(constants.SECRETE_JWT), nil // Replace with your actual secret key
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
