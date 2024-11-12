package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sudarshan284/gin-app/controller"
	"github.com/sudarshan284/gin-app/database"
	"github.com/sudarshan284/gin-app/models"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	loadEnv()
	loadDatabase()
	serveApp()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}

func serveApp() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	router.Run(":8001")
	fmt.Println("Server running on port 8001")
}
