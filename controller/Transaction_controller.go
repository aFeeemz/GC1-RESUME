package controller

import (
	"GC1/models"
	"GC1/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	TR repository.TransactionRepo
}

func (c *TransactionController) CreateTransaction(e echo.Context) error {
	// Body request
	var req models.Transaction
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("INPUT DATA CREATE TRANSACTION :", req)

	if req.FieldID == primitive.NilObjectID || req.UserID == primitive.NilObjectID {
		return e.JSON(http.StatusBadRequest, "error params")
	}

	// Create transactions
	result, err := c.TR.CreateTransaction(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, result)
}
