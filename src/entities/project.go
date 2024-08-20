package entities

import (
	"time"
)

type Project struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"notNull"`
	Url       *string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

func (Project) TableName() string {
	return "master_projects"
}
