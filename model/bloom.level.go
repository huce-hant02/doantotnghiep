package model

import (
	"time"

	"gorm.io/gorm"
)

// Cấp độ
type Level struct {
	ID           uint        `json:"id"`
	BloomGroupID uint        `json:"bloomGroupId" gorm:"uniqueIndex:bloom_level_uniq, where deleted_at is null"`
	Group        *BloomGroup `json:"group" gorm:"foreignKey:BloomGroupID;constraint:OnUpdate:CASCADE;"`
	Active       *bool       `json:"active" gorm:"default:true"`

	Title       string    `json:"title" gorm:"uniqueIndex:bloom_level_uniq, where deleted_at is null"`
	Level       int       `json:"level"`
	Description string    `json:"description"`
	Keywords    []Keyword `json:"keywords" gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
