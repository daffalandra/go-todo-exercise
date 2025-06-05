package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/daffalandra/go-todo-exercise/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global GORM database connection
var DB *gorm.DB

// TemplateDir is the base directory for templates
var TemplateDir = "views"

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// InitDB initializes the database connection with GORM
func InitDB() (*gorm.DB, error) {
	dsn, err := GetDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Auto-migrate the schema
	err = db.AutoMigrate(&entities.Category{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate schema: %w", err)
	}

	DB = db
	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	return sqlDB.Close()
}

// GetDSN constructs the Data Source Name
func GetDSN() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || dbname == "" {
		return "", fmt.Errorf("missing required environment variables")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port), nil
}

// GetTemplatePath returns the full path to a template file
func GetTemplatePath(templateName string) string {
	return filepath.Join(TemplateDir, templateName)
}
