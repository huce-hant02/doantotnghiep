package model

import (
	"time"
)

type OutlineDocument struct {
	ID uint `json:"id" gorm:"unique;autoIncrement"`

	OutlineID  uint `json:"outlineId"`
	DocumentID uint `json:"documentId"`

	Type     string `json:"type"`
	Indexing int    `json:"indexing"`

	Outline  *Outline  `json:"outline" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Document *Document `json:"document" gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
