package model

import (
	"time"
)

// ModelType định nghĩa model
type ModelType struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	Code   string            `json:"code" gorm:"uniqueIndex:model_type_code_uniq"`
	Name   string            `json:"name"`
	Fields *[]ModelTypeField `json:"fields" gorm:"foreignKey:ModelTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Active *bool             `json:"active" gorm:"default:true"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	OtherFilter OtherRFPFilter `json:"otherFilter" gorm:"-"`
}
