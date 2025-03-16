package model

import (
	"time"
)

type SchoolYearEducationProgram struct {
	ID                 uint `json:"id" gorm:"unique;autoIncrement"`
	SchoolYearID       uint `json:"schoolYearId"`
	EducationProgramID uint `json:"educationProgramId"`
	OrdinalNumber      int  `json:"ordinalNumber"` // hoc ky thu 1, hoc ky thu 2

	SchoolYear       *SchoolYear       `json:"schoolYear" gorm:"foreignKey:SchoolYearID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationProgram *EducationProgram `json:"educationProgram" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
