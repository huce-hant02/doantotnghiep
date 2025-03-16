package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ResultEvaluate struct {
	ID                uint           `json:"id"`
	OutlineID         uint           `json:"outlineId"`
	CoursePointWeight float32        `json:"coursePointWeight"`
	PostCode          string         `json:"postCode"`
	PostTitle         string         `json:"postTitle"`
	EvaluateCriteria  pq.StringArray `json:"evaluateCriteria" gorm:"type:text[];default:'{}'"`
	MaxPoint          float32        `json:"maxPoint"`
	Type              string         `json:"type"`
	ExamTime          int            `json:"examTime"`

	Outline *Outline `json:"-" swaggerignore:"true" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	// foreign key data
	EvaluateForms []ResultEvaluateExamType `json:"evaluateForms" gorm:"foreignKey:ResultEvaluateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rubrics       []ResultEvaluateRubric   `json:"rubrics" gorm:"foreignKey:ResultEvaluateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostOutputs   []PostOutput             `json:"postOutputs" gorm:"foreignKey:ResultEvaluateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ResultEvaluateRepository interface {
	Insert(*gorm.DB, []ResultEvaluate) ([]ResultEvaluate, error)
	Update(*gorm.DB, *ResultEvaluate) (*ResultEvaluate, error)
	Delete(db *gorm.DB, ID uint) error
	IAssociationOpRepo
	ISearchRepo
}
