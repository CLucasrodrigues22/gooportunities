package main

import (
	"github.com/CLucasrodrigues22/gooportunities/config"
	"github.com/CLucasrodrigues22/gooportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization failed: %v", err)
		return
	}

	// Initialize Router
	router.InitializeRouter()
}
