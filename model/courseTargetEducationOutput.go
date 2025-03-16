package model

import (
	"time"
)

type CourseTargetEducationOutput struct {
	ID                        uint `json:"id" gorm:"autoIncrement"`
	CourseTargetID            uint `json:"courseTargetId"`
	EducationStandardOutputID uint `json:"educationStandardOutputId"`

	CourseTarget            *CourseTarget            `json:"courseTarget" swaggerignore:"true" gorm:"foreignKey:CourseTargetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationStandardOutput *EducationStandardOutput `json:"educationStandardOutput" swaggerignore:"true" gorm:"foreignKey:EducationStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}
