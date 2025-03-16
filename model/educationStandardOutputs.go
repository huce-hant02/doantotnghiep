package model

import (
	"time"

	"gorm.io/gorm"
)

type EducationStandardOutput struct {
	ID                 uint   `json:"id"`
	EducationProgramID uint   `json:"educationProgramId" gorm:"uniqueIndex:education_output_uniq, where deleted_at is null"`
	Indexing           string `json:"indexing" gorm:"uniqueIndex:education_output_uniq, where deleted_at is null"`

	Description       string  `json:"description"`
	AbilityLevel      float32 `json:"abilityLevel"`
	OutputLevel       int     `json:"outputLevel"`
	EducationTargetID uint    `json:"educationTargetId"`

	KeywordID uint           `json:"keywordId"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	Keyword *Keyword `json:"keyword" swaggerignore:"true" gorm:"foreignKey:KeywordID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`

	// foreignKey
	EducationTarget *EducationTarget `json:"educationTarget" gorm:"foreignKey:EducationTargetID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	ITUTables       []ITUTable       `swaggerignore:"true" json:"-" gorm:"foreignKey:EducationStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
