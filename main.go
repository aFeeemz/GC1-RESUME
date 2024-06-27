package main

import (
	"GC1/config"
	"GC1/models"
	"GC1/router"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, db := config.Connect(context.Background(), "rentfield_gc")
	defer client.Disconnect(context.Background())

	e := echo.New()

	router.InitRoutes(e, db)
	// pushdata(db) //push fields data once
	// get port from .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

}

func pushdata(db *mongo.Database) {
	fields := []interface{}{
		models.Field{Name: "Field 1", Price: 55.0, Status: "Available"},
		models.Field{Name: "Field 2", Price: 66.0, Status: "Available"},
		models.Field{Name: "Field 3", Price: 75.0, Status: "Unavailable"},
	}

	result, err := db.Collection("fields").InsertMany(context.TODO(), fields)
	if err != nil {
		return
	}

	fmt.Println("Fields :", result)
}
