package model

import (
	"time"

	"github.com/lib/pq"
)

type TeachingPlan struct {
	ID               uint           `json:"id"`
	OutlineID        uint           `json:"outlineId"`
	Index            string         `json:"index"`
	NumberOfLessons  *float32       `json:"numberOfLessons" gorm:"default:0"`
	ContentType      string         `json:"contentType"`
	Content          string         `json:"content"`
	TeachingMethod   pq.StringArray `json:"teachingMethod" gorm:"type:text[];default:'{}'"`
	TeachingActivity string         `json:"teachingActivity"`
	StudyingActivity string         `json:"studyingActivity"`
	// AssessmentPost   pq.StringArray `json:"assessmentPost" gorm:"type:text[];default:'{}'"`

	NumOfLtPeriod    *float32 `json:"numOfLtPeriod"`
	NumOfThPeriod    *float32 `json:"numOfThPeriod"`
	SelfTaughtPeriod *float32 `json:"selfTaughtPeriod"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	// foreign key
	StandardOutputs []TeachingPlanStandardOutput `json:"standardOutputs" gorm:"foreignKey:TeachingPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AssessmentPosts []TeachingPlanResultEvaluate `json:"assessmentPosts" gorm:"foreignKey:TeachingPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Documents       []TeachingPlanDocument       `json:"documents" gorm:"foreignKey:TeachingPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TeachingPlanRepository interface {
}
