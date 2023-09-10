package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	routes "github.com/piyushbd14/golang-jwt-project/routes"

	"github.com/gin-gonic/gin" //define routes for handling HTTP requests        Gin supports middleware, which allows you to inject functionality into the request/response chain. Middleware can be used for tasks like authentication, logging, request/response modification, and more.              Gin includes built-in support for parsing JSON and XML request bodies and rendering JSON and XML responses
)

func main() {
	err := godotenv.Load(".env") //to load .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": ""})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "success granted for api-2"})
	})
	router.Run(":" + port)
}
