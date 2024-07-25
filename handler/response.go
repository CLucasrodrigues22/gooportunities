package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(context *gin.Context, code int, msg string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"massage": msg,
		"error":   code,
	})
}

func sendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}
