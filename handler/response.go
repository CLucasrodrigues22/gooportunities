package handler

import (
	"fmt"
	"github.com/CLucasrodrigues22/gooportunities/schemas"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
