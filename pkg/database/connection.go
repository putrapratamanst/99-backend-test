package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite"
)

type Connection struct {
	DB *gorm.DB
}
type DatabaseConfig struct {
	DBPath string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBPath: getEnv("DB_PATH", "./database.db"),
	}
}
func Connect(config *DatabaseConfig) (*Connection, error) {
	sqlDB, err := sql.Open("sqlite", config.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Printf("Successfully connected to SQLite database at: %s", config.DBPath)
	return &Connection{DB: db}, nil
}
func (c *Connection) AutoMigrate(models ...interface{}) error {
	return c.DB.AutoMigrate(models...)
}
func (c *Connection) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
