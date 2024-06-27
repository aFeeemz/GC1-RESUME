package router

import (
	"GC1/controller"
	"GC1/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, db *mongo.Database) {
	e.Use(middleware.Recover())

	repo := &repository.Repo{DB: db}

	// transaction controller init
	tc := &controller.TransactionController{TR: repo}
	uc := &controller.UserController{UR: repo}
	fc := &controller.FieldController{FR: repo}

	e.POST("/Transaction", tc.CreateTransaction)
	e.POST("/Create-Users", uc.CreateUser)

	e.GET("/All-Fields", fc.GetAllFields)

}
