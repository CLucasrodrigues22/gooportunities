package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(pfx string) *Logger {
	// Initialize Logger
	logger = NewLogger(pfx)
	return logger
}
