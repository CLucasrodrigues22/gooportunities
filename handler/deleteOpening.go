package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening",
	})
}
