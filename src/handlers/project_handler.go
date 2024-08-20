package handlers

import (
	"net/http"
	"simple-gin/src/dto"
	errorhandlers "simple-gin/src/error-handlers"
	"simple-gin/src/helpers"
	"simple-gin/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type projectHandler struct {
	service services.ProjectServiceType
}

func ProjectHandler(service services.ProjectServiceType) *projectHandler {
	return &projectHandler{
		service: service,
	}
}

func (handler *projectHandler) IndexProjects(context *gin.Context) {
	data, err := handler.service.IndexProjects()

	if err != nil {
		errorhandlers.HandleError(context, err)

		return
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Get index projects successfully",
		Data:       data,
	})

	context.JSON(http.StatusOK, response)
}

func (handler *projectHandler) DetailProject(context *gin.Context) {
	projectId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		errorhandlers.HandleError(context, &errorhandlers.BadRequestError{
			Message: "Id must not null and integer",
		})

		return
	}

	data, err := handler.service.DetailProject(projectId)

	if err != nil {
		errorhandlers.HandleError(context, err)

		return
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Get detail project successfully",
		Data:       data,
	})

	context.JSON(http.StatusOK, response)
}

func (handler *projectHandler) CreateProject(context *gin.Context) {
	var request dto.StoreProjectRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		errorhandlers.HandleError(context, &errorhandlers.BadRequestError{
			Message: err.Error(),
		})

		return
	}

	if err := handler.service.CreateProject(&request); err != nil {
		errorhandlers.HandleError(context, err)

		return
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create project successfully",
	})

	context.JSON(http.StatusCreated, response)
}

func (handler *projectHandler) UpdateProject(context *gin.Context) {
	var request dto.UpdateProjectRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		errorhandlers.HandleError(context, &errorhandlers.BadRequestError{
			Message: err.Error(),
		})

		return
	}

	projectId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		errorhandlers.HandleError(context, &errorhandlers.BadRequestError{
			Message: "Id must not null and integer",
		})

		return
	}

	if err := handler.service.UpdateProject(projectId, &request); err != nil {
		errorhandlers.HandleError(context, err)

		return
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update project successfully",
	})

	context.JSON(http.StatusOK, response)
}

func (handler *projectHandler) DeleteProject(context *gin.Context) {
	projectId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		errorhandlers.HandleError(context, &errorhandlers.BadRequestError{
			Message: "Id must not null and integer",
		})

		return
	}

	if err := handler.service.DeleteProject(projectId); err != nil {
		errorhandlers.HandleError(context, err)

		return
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Delete project successfully",
	})

	context.JSON(http.StatusOK, response)
}
