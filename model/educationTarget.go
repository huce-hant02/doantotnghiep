package model

import (
	"time"

	"gorm.io/gorm"
)

type EducationTarget struct {
	ID                 uint   `json:"id"`
	EducationProgramID uint   `json:"educationProgramId" gorm:"uniqueIndex:education_target_uniq, where deleted_at is null"`
	Code               string `json:"code" gorm:"uniqueIndex:education_target_uniq, where deleted_at is null"`

	Title   string `json:"title"`
	Content string `json:"content"`

	SortIndex int `json:"sortIndex"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
