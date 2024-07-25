package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize MYSQL
	db, err = InitializeMySQL()
	if err != nil {
		return fmt.Errorf("init mysql err: %v", err)
	}
	return nil
}

func GetMySQL() *gorm.DB {
	return db
}

func GetLogger(pfx string) *Logger {
	// Initialize Logger
	logger = NewLogger(pfx)
	return logger
}
