package model

import (
	"time"
)

// ModelTypeField định nghĩa model và attributes
type ModelTypeField struct {
	ID          uint `json:"id" gorm:"autoIncrement"`
	ModelTypeId uint `json:"modelTypeId" gorm:"uniqueIndex:model_type_field_code_uniq"`

	Code   string `json:"code" gorm:"uniqueIndex:model_type_field_code_uniq"`
	Name   string `json:"name"`
	Type   string `json:"type"` // text | float | int | ... | preload
	Active *bool  `json:"active" gorm:"default:true"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	// FK
	ModelType *ModelType `swaggerignore:"true" json:"-" gorm:"foreignKey:ModelTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
