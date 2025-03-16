package model

import (
	"time"

	"github.com/lib/pq"
)

type CourseStandardOutput struct {
	ID        uint `json:"id"`
	OutlineID uint `json:"outlineId"`
	// EducationStandardOutputCode string         `json:"educationStandardOutputCode"`
	Indexing      string         `json:"indexing"`
	Description   string         `json:"description"`
	TeachingLevel pq.StringArray `json:"teachingLevel" gorm:"type:text[];default:'{}'"`
	KeywordID     uint           `json:"keywordId"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	Keyword                  Keyword                       `json:"keyword" swaggerignore:"true" gorm:"foreignKey:KeywordID"`
	EducationStandardOutputs []CourseOutputEducationOutput `json:"educationStandardOutputs" gorm:"foreignKey:CourseStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
