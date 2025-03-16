package model

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name" gorm:"unique"`
	FacultyID uint           `json:"facultyId"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type SubjectRepository interface {
	IBasicDBQueryRepo
	ISearchRepo
}
