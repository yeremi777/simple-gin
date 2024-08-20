package repositories

import (
	"simple-gin/src/entities"

	"gorm.io/gorm"
)

type ProjectRepositoryType interface {
	IndexProjects() ([]entities.Project, error)
	DetailProject(id int) (entities.Project, error)
	CreateProject(project *entities.Project) error
	UpdateProject(project *entities.Project, updateData *map[string]interface{}) error
	DeleteProject(id int) error
}

type projectRepository struct {
	database *gorm.DB
}

func ProjectRepository(database *gorm.DB) *projectRepository {
	return &projectRepository{
		database: database,
	}
}

func (repository *projectRepository) IndexProjects() ([]entities.Project, error) {
	var projects []entities.Project

	err := repository.database.Find(&projects).Error

	return projects, err
}

func (repository *projectRepository) DetailProject(id int) (entities.Project, error) {
	var project entities.Project = entities.Project{ID: uint(id)}

	err := repository.database.First(&project).Error

	return project, err
}

func (repository *projectRepository) CreateProject(project *entities.Project) error {
	err := repository.database.Create(&project).Error

	return err
}

func (repository *projectRepository) UpdateProject(project *entities.Project, updateData *map[string]interface{}) error {
	err := repository.database.Model(&project).Updates(&updateData).Error

	return err
}

func (repository *projectRepository) DeleteProject(id int) error {
	err := repository.database.Delete(&entities.Project{}, id).Error

	return err
}
