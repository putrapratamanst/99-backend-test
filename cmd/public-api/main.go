package main

import (
	"99-backend-exercise/internal/publicapi"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://localhost:8001"
	}
	listingServiceURL := os.Getenv("LISTING_SERVICE_URL")
	if listingServiceURL == "" {
		listingServiceURL = "http://localhost:6000"
	}
	serviceClient := publicapi.NewServiceClient(userServiceURL, listingServiceURL)
	publicAPIService := publicapi.NewService(serviceClient)
	publicAPIHandler := publicapi.NewHandler(publicAPIService)
	router := gin.Default()
	publicAPIGroup := router.Group("/public-api")
	{
		publicAPIGroup.GET("/listings", publicAPIHandler.GetListings)
		publicAPIGroup.POST("/users", publicAPIHandler.CreateUser)
		publicAPIGroup.POST("/listings", publicAPIHandler.CreateListing)
	}
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "public-api"})
	})
	port := os.Getenv("PUBLIC_API_PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Public API service starting on port %s", port)
	log.Printf("User Service URL: %s", userServiceURL)
	log.Printf("Listing Service URL: %s", listingServiceURL)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
