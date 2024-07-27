package config

import (
	"fmt"
	"github.com/CLucasrodrigues22/gooportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	// Load .env credentials
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("File .env not exist: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	// Create DB and connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql connect err: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("mysql migrate err: %v", err)
		return nil, err
	}

	return db, nil
}
