package model

import (
	"time"

	"gorm.io/gorm"
)

// Bloom - Nhóm
type BloomGroup struct {
	ID          uint       `json:"id"`
	BloomWordID uint       `json:"bloomWordId" gorm:"uniqueIndex:bloom_group_uniq, where deleted_at is null"`
	Area        *BloomWord `json:"area" gorm:"foreignKey:BloomWordID;constraint:OnUpdate:CASCADE;"`
	Active      *bool      `json:"active" gorm:"default:true"`

	Code        string  `json:"code" gorm:"uniqueIndex:bloom_group_uniq, where deleted_at is null"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Levels      []Level `json:"levels" gorm:"foreignKey:BloomGroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
