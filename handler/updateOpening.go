package handler

import (
	"github.com/CLucasrodrigues22/gooportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validate UpdateOpeningRequest error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, "id can't be empty")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}
	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary <= 0 {
		opening.Salary = request.Salary
	}
	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("update opening error: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "update opening error")
		return
	}
	sendSuccess(context, "update-opening", opening)
}
