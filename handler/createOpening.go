package handler

import (
	"github.com/CLucasrodrigues22/gooportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("bind json error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validate error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("DB Error: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "Error creating opening.")
		return
	}

	sendSuccess(context, "create-opening", opening)
}
