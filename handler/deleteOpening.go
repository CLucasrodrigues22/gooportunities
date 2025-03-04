package handler

import (
	"fmt"
	"github.com/CLucasrodrigues22/gooportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("Opening with id: %s not found", id))
		return
	}
	// delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("Deleting opening with id: %s failed", id))
		return
	}
	sendSuccess(context, "delete-opening", opening)
}
