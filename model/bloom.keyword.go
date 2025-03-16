package model

import (
	"time"

	"gorm.io/gorm"
)

// Từ khoá
type Keyword struct {
	ID      uint   `json:"id"`
	LevelID uint   `json:"levelId" gorm:"uniqueIndex:bloom_keyword_uniq, where deleted_at is null"`
	Value   string `json:"value" gorm:"uniqueIndex:bloom_keyword_uniq, where deleted_at is null"`
	Active  *bool  `json:"active" gorm:"default:true"`

	Level *Level `json:"level" swaggerignore:"true" gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
