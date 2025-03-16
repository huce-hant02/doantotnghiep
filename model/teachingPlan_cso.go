package model

import "time"

type TeachingPlanStandardOutput struct {
	ID                     uint `json:"id" gorm:"unique;autoIncrement"`
	TeachingPlanID         uint `json:"teachingPlanId"`
	CourseStandardOutputID uint `json:"courseStandardOutputId"`

	CourseStandardOutput *CourseStandardOutput `swaggerignore:"true" json:"courseStandardOutput" gorm:"foreignKey:CourseStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeachingPlan         *TeachingPlan         `swaggerignore:"true" json:"teachingPlan" gorm:"foreignKey:TeachingPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}
