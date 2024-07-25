package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET openings",
	})
}
