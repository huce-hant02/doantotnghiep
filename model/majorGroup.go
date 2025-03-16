package model

import (
	"time"

	"gorm.io/gorm"
)

type MajorGroup struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:major_group_unique"`
	Code                 string `json:"code" gorm:"uniqueIndex:major_group_unique"`

	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type MajorGroupRepository interface {
	ISearchRepo
}
