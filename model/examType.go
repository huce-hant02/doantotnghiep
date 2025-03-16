package model

import (
	"time"

	"gorm.io/gorm"
)

// Hình thức thi
type ExamType struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:exam_type_uniq"`
	Code                 string `json:"code" gorm:"uniqueIndex:exam_type_uniq"`
	Active               *bool  `json:"active" gorm:"default:true"`

	Title           string             `json:"title"`
	HasRubric       *bool              `json:"hasRubric" gorm:"default:false"`
	RequireExamTime *bool              `json:"requireExamTime" gorm:"default:false"`
	TemplateRubrics ListTemplateRubric `json:"templateRubrics" gorm:"type:text"` // Rubric mẫu

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
