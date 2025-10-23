package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

// Connection holds the database connection
type Connection struct {
	DB *gorm.DB
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	DBPath string
}

// NewDatabaseConfig creates a new database config from environment variables
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBPath: getEnv("DB_PATH", "./database.db"),
	}
}

// Connect establishes a connection to the database
func Connect(config *DatabaseConfig) (*Connection, error) {
	// Open SQLite connection using pure Go driver
	sqlDB, err := sql.Open("sqlite", config.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	// Use the existing sql.DB connection with GORM
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Successfully connected to SQLite database at: %s", config.DBPath)
	return &Connection{DB: db}, nil
}

// AutoMigrate runs database migrations
func (c *Connection) AutoMigrate(models ...interface{}) error {
	return c.DB.AutoMigrate(models...)
}

// Close closes the database connection
func (c *Connection) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
