package entities

import (
	"time"

	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"notNull"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}
