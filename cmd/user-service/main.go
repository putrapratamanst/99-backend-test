package main

import (
	"99-backend-exercise/internal/models"
	"99-backend-exercise/internal/user"
	"99-backend-exercise/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Setup database connection
	dbConfig := database.NewDatabaseConfig()
	dbConn, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	// Run migrations
	if err := dbConn.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize dependencies
	userRepo := user.NewRepository(dbConn.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// Setup router
	router := gin.Default()

	// User routes
	v1 := router.Group("/")
	{
		v1.GET("/users", userHandler.GetUsers)
		v1.GET("/users/:id", userHandler.GetUserByID)
		v1.POST("/users", userHandler.CreateUser)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "user-service"})
	})

	// Get port from environment or use default
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8001"
	}

	log.Printf("User service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
