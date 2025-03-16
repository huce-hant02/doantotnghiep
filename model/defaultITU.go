package model

import (
	"time"

	"gorm.io/gorm"
)

type DefaultITU struct {
	ID                   uint `json:"id" gorm:"uniqueIndex:scholastic_default_itu_uniq, where deleted_at is null"`
	ScholasticSemesterID uint `json:"scholasticSemesterId" gorm:"uniqueIndex:scholastic_default_itu_uniq, where deleted_at is null"`

	CourseID                uint `json:"courseId"`
	DefaultStandardOutputID uint `json:"defaultStandardOutputId"`
	// mức độ
	Level string `json:"level"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

type DefaultITURepository interface {
	ISearchRepo
}
