package controller

import (
	"GC1/repository"
	"GC1/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FieldController struct {
	FR repository.FieldRepo
}

func (c *FieldController) GetAllFields(e echo.Context) error {
	result, err := c.FR.GetAllFields()
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error,
			"message": utils.ErrGetData,
		})
	}

	return e.JSON(http.StatusOK, result)

}
