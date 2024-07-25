package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "PUT opening",
	})
}
