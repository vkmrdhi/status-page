package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
	Auth0Domain   string
	Auth0ClientID string
	Auth0Audience string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		Auth0Domain:   os.Getenv("AUTH0_DOMAIN"),
		Auth0ClientID: os.Getenv("AUTH0_CLIENT_ID"),
		Auth0Audience: os.Getenv("AUTH0_AUDIENCE"),
	}
}

func InitializeDatabase(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Add any additional GORM configurations
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func MigrateDatabase(db *gorm.DB) error {
	// Migrate all models
	return db.AutoMigrate(
		&models.User{},
		&models.Organization{},
		&models.Service{},
		&models.Incident{},
		&models.Maintenance{},
	)
}
