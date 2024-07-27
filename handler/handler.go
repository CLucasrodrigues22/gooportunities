package handler

import (
	"github.com/CLucasrodrigues22/gooportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySQL()
}
