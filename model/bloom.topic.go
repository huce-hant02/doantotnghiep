package model

import (
	"time"

	"gorm.io/gorm"
)

// Chủ đề
type BloomWord struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:bloom_topic_uniq, where deleted_at is null"`
	Code                 string `json:"code" gorm:"uniqueIndex:bloom_topic_uniq, where deleted_at is null"`
	Active               *bool  `json:"active" gorm:"default:true"`

	Title       string       `json:"title"`
	Description string       `json:"description"`
	Groups      []BloomGroup `json:"groups" gorm:"foreignKey:BloomWordID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
