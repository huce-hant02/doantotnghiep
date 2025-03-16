package model

import (
	"time"

	"github.com/lib/pq"
)

type CourseTarget struct {
	ID           uint           `json:"id"`
	OutlineID    uint           `json:"outlineId"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	AbilityLevel pq.StringArray `json:"abilityLevel" gorm:"type:text[];default:'{}'"`
	KeywordID    uint           `json:"keywordId"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	Keyword                  Keyword                       `json:"keyword" swaggerignore:"true" gorm:"foreignKey:KeywordID"`
	EducationStandardOutputs []CourseTargetEducationOutput `json:"educationStandardOutputs" gorm:"foreignKey:CourseTargetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
