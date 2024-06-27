package controller

import (
	"GC1/models"
	"GC1/repository"
	"GC1/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UR repository.UserRepo
}

func (c *UserController) CreateUser(e echo.Context) error {
	// Body request
	var req models.User
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Create user
	result, err := c.UR.CreateUser(&req)
	if err != nil {
		if err == utils.ErrDuplicateUser {
			return e.JSON(http.StatusConflict, echo.Map{
				"error": err.Error(),
			})
		}
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("INPUT DATA CREATE TRANSACTION :", req)
	return e.JSON(http.StatusCreated, result)
}
