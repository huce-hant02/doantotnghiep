package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseType struct {
	ID                   uint `json:"id"`
	ScholasticSemesterID uint `json:"scholasticSemesterId" gorm:"uniqueIndex:course_type_uniq, where deleted_at is null"`

	Active              *bool    `json:"active" gorm:"default:true"`
	Code                string   `json:"code" gorm:"uniqueIndex:course_type_uniq, where deleted_at is null"`
	Title               string   `json:"title"`
	EvaluateForm        int      `json:"evaluateForm"`
	RequireExamTime     *bool    `json:"requireExamTime"`
	NumOfCreditClinical *float32 `json:"numOfCreditClinical" gorm:"default:0"` // STC lâm sàng

	// OutlineDeadline       *time.Time `json:"outlineDeadline"`
	MaxTeacherInCourse    *int `json:"maxTeacherInCourse" gorm:"default:0"`    // Số giảng viên tối đa cho mỗi học phần
	MinTeacherInCourse    *int `json:"minTeacherInCourse" gorm:"default:0"`    // Số giảng viên tối thiểu cho mỗi học phần
	MaxCriteriaInClinical *int `json:"maxCriteriaInClinical" gorm:"default:0"` // Số tiêu chí đánh giá tối đa cho mỗi bảng đánh giá lâm sàng
	MinCriteriaInClinical *int `json:"minCriteriaInClinical" gorm:"default:0"` // Số tiêu chí đánh giá tối thiểu cho mỗi bảng đánh giá lâm sàng
	MaxCdioOutcomes       *int `json:"maxCdioOutcomes" gorm:"default:0"`       // CDIO - Số CĐR học phần tối đa
	MinCdioOutcomes       *int `json:"minCdioOutcomes" gorm:"default:0"`       // CDIO - Số CĐR học phần tối thiểu
	MaxAbetOutcomes       *int `json:"maxAbetOutcomes" gorm:"default:0"`       // ABET - Số CĐR học phần tối đa
	MinAbetOutcomes       *int `json:"minAbetOutcomes" gorm:"default:0"`       // ABET - Số CĐR học phần tối thiểu

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

type CourseTypeRepository interface {
	Create(db *gorm.DB, CourseTypes []CourseType, listOmitField []string) ([]CourseType, error)
	Update(db *gorm.DB, ID uint, CourseType *CourseType) (*CourseType, error)
	Delete(db *gorm.DB, ID uint) error
	ISearchRepo
}
