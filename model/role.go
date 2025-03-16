package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	Code        string         `json:"code" gorm:"unique"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Active      *bool          `json:"active" gorm:"default:true"`
	Acronym     pq.StringArray `json:"acronym" gorm:"type:text"` // Ký hiệu position-detail-code từ HRM

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type RoleRepo interface {
}
