package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {

	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

}

func CreateUser(c echo.Context) (err error) {
	var user User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, user)

}
