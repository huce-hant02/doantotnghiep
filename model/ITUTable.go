package model

import (
	"time"
)

type ITUTable struct {
	ID       uint `json:"id"`
	CourseID uint `json:"courseId"  gorm:"index:itu_table_uniq,unique"`
	// CourseEducationProgramID  uint `json:"courseEducationProgramId"  gorm:"index:itu_table_uniq,unique"`
	EducationStandardOutputID uint `json:"educationStandardOutputId"  gorm:"index:itu_table_uniq,unique"`
	EducationProgramID        uint `json:"educationProgramId"`
	// mức độ
	Level string `json:"level"`

	// EducationStandardOutput *EducationStandardOutput `json:"educationStandardOutput" gorm:"foreignKey:EducationStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// CourseEducationProgram  *CourseEducationProgram  `json:"courseEducationProgram" gorm:"foreignKey:CourseEducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
