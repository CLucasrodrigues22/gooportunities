package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET opening",
			})
		})
		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "POST opening",
			})
		})
		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "DELETE opening",
			})
		})
		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "PUT opening",
			})
		})
		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET openings",
			})
		})
	}
}
