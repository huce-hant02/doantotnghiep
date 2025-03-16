package model

import (
	"time"

	"gorm.io/gorm"
)

type Scholastic struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	Status      string `json:"status"`
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`

	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`

	ScholasticSemesters []ScholasticSemester `json:"scholasticSemesters" gorm:"foreignKey:ScholasticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ScholasticRepository interface {
	ISearchRepo
}
