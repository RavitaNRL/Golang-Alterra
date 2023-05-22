package usecase

import (
	"code_competence/middleware"
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(req *payload.CreateUserRequest) (*model.User, error) {
	_, err := database.GetUserByEmail(req.Email)
	if err == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Already Registered")
	}
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = database.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(email, password string) (*model.User, error) {
	user, err := database.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Not Registered")
	}
	if user.Password != password {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid Password")
	}

	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Error Create Token")
	}
	user.Token = token
	return user, nil
}
