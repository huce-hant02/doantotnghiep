package model

import (
	"time"

	"gorm.io/gorm"
)

type Specialty struct {
	ID        uint           `json:"id"`
	Code      string         `json:"code"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
type SpecialtyRepository interface {
	ISearchRepo
}
