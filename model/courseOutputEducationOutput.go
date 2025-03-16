package model

import (
	"time"
)

type CourseOutputEducationOutput struct {
	ID                        uint `json:"id" gorm:"autoIncrement"`
	CourseStandardOutputID    uint `json:"courseStandardOutputId"`
	EducationStandardOutputID uint `json:"educationStandardOutputId"`

	CourseStandardOutput    *CourseStandardOutput    `json:"courseStandardOutput" swaggerignore:"true" gorm:"foreignKey:CourseStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationStandardOutput *EducationStandardOutput `json:"educationStandardOutput" swaggerignore:"true" gorm:"foreignKey:EducationStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}
