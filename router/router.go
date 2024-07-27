package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeRouter() {
	// Initialize router
	router := gin.Default()
	// Initialize routes
	initializeRoutes(router)
	// Run the server
	routePort := os.Getenv("APP_PORT")
	router.Run(":" + routePort) // listen and serve on 0.0.0.0:APP_PORT
}
