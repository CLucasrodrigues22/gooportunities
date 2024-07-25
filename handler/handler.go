package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "POST opening",
	})
}
func ShowOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "PUT opening",
	})
}
func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening",
	})
}
func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET openings",
	})
}
