package model

import "time"

type PostOutput struct {
	ID                      uint `json:"id"`
	ResultEvaluateID        uint `json:"resultEvaluateId" gorm:"uniqueIndex:result_evaluate_output_uniq"`
	EvaluatedStandardOutput uint `json:"evaluatedStandardOutput" gorm:"uniqueIndex:result_evaluate_output_uniq"`

	MaxPoint                  float32 `json:"maxPoint"`                  // '3'
	StandardOutputPointWeight float32 `json:"standardOutputPointWeight"` // `1`

	CourseStandardOutput *CourseStandardOutput `json:"-" swaggerignore:"true" gorm:"foreignKey:EvaluatedStandardOutput;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}
