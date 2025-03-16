package model

import (
	"time"

	"gorm.io/gorm"
)

type SchoolYear struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	StartAt       string         `json:"startAt"`
	OrdinalNumber int            `json:"-" gorm:"-"`
	CreatedAt     time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt     gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt     time.Time      `json:"updatedAt" swaggerignore:"true"`
}
