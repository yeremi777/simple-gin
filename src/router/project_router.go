package router

import (
	"simple-gin/configs"
	"simple-gin/src/handlers"
	"simple-gin/src/repositories"
	"simple-gin/src/services"

	"github.com/gin-gonic/gin"
)

func ProjectRouter(route *gin.RouterGroup) {
	projectRepository := repositories.ProjectRepository(configs.Database)
	projectService := services.ProjectService(projectRepository)
	projectHandler := handlers.ProjectHandler(projectService)

	route.GET("/", projectHandler.IndexProjects)
	route.GET("/:id", projectHandler.DetailProject)
	route.POST("", projectHandler.CreateProject)
	route.PUT("/:id", projectHandler.UpdateProject)
	route.DELETE("/:id", projectHandler.DeleteProject)
}
