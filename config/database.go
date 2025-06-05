package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig holds PostgreSQL connection details
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var Config DBConfig

// LoadEnv loads environment variables and stores them in Config
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Config = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

// GetDSN returns the PostgreSQL DSN string
func GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Config.Host, Config.User, Config.Password, Config.DBName, Config.Port,
	)
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
