package model

import "time"

type ResultEvaluateRubric struct {
	ID               uint `json:"id"`
	ResultEvaluateID uint `json:"resultEvaluateId" gorm:"uniqueIndex:result_evaluate_rubric_uniq"`
	RubricID         uint `json:"rubricId" gorm:"uniqueIndex:result_evaluate_rubric_uniq"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	Rubric *Rubric `json:"rubric" swaggerignore:"true" gorm:"foreignKey:RubricID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
