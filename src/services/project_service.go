package services

import (
	"fmt"
	"simple-gin/src/dto"
	"simple-gin/src/entities"
	errorhandlers "simple-gin/src/error-handlers"
	"simple-gin/src/repositories"
)

type ProjectServiceType interface {
	IndexProjects() (*[]dto.Project, error)
	DetailProject(id int) (*dto.Project, error)
	CreateProject(request *dto.StoreProjectRequest) error
	UpdateProject(id int, request *dto.UpdateProjectRequest) error
	DeleteProject(id int) error
}

type projectService struct {
	repository repositories.ProjectRepositoryType
}

func ProjectService(repository repositories.ProjectRepositoryType) *projectService {
	return &projectService{
		repository: repository,
	}
}

func (service *projectService) IndexProjects() (*[]dto.Project, error) {
	data := make([]dto.Project, 0)

	projects, err := service.repository.IndexProjects()

	if err != nil {
		return nil, &errorhandlers.InternalServerError{
			Message: err.Error(),
		}
	}

	for _, project := range projects {
		pj := dto.Project{
			ID:   project.ID,
			Name: project.Name,
			URL:  *project.Url,
		}

		data = append(data, pj)
	}

	return &data, nil
}

func (service *projectService) DetailProject(id int) (*dto.Project, error) {
	project, err := service.repository.DetailProject(id)

	if err != nil {
		message := fmt.Sprintf("Project with id %v not found", id)

		return nil, &errorhandlers.NotFoundError{
			Message: message,
		}
	}

	return &dto.Project{
		ID:   project.ID,
		Name: project.Name,
		URL:  *project.Url,
	}, nil
}

func (service *projectService) CreateProject(request *dto.StoreProjectRequest) error {
	if request.Name == "" {
		return &errorhandlers.BadRequestError{
			Message: "Name can't be null",
		}
	}

	project := entities.Project{
		Name: request.Name,
		Url:  &request.URL,
	}

	if err := service.repository.CreateProject(&project); err != nil {
		return &errorhandlers.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (service *projectService) UpdateProject(id int, request *dto.UpdateProjectRequest) error {
	project, err := service.repository.DetailProject(id)

	if err != nil {
		message := fmt.Sprintf("Project with id %v not found", id)

		return &errorhandlers.NotFoundError{
			Message: message,
		}
	}

	updatedData := make(map[string]interface{})

	if request.Name != "" {
		updatedData["name"] = request.Name
	}

	if request.URL != "" {
		updatedData["url"] = request.URL
	}

	if err := service.repository.UpdateProject(&project, &updatedData); err != nil {
		return &errorhandlers.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (service *projectService) DeleteProject(id int) error {
	if err := service.repository.DeleteProject(id); err != nil {
		return &errorhandlers.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
