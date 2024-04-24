package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/model"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serverApplication() {
	router := gin.Default()

	publicRouters := router.Group("/auth")
	publicRouters.POST("/register", controller.Register)
	publicRouters.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	port := "8080"
	if os.Getenv("PORT") == "" {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
	fmt.Println("Server running on port 8000")
}
