package model

import "time"

type ResultEvaluateExamType struct {
	ID               uint `json:"id"`
	ResultEvaluateID uint `json:"resultEvaluateId" gorm:"uniqueIndex:result_evaluate_exam_type_uniq"`
	ExamTypeID       uint `json:"examTypeId" gorm:"uniqueIndex:result_evaluate_exam_type_uniq"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	ExamType *ExamType `json:"examType" swaggerignore:"true" gorm:"foreignKey:ExamTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
