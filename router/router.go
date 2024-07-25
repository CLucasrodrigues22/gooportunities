package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Initialize router
	router := gin.Default()
	// Initialize routes
	initializeRoutes(router)
	// Run the server
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
