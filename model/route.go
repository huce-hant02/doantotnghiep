package model

import (
	"time"

	"gorm.io/gorm"
)

type Route struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	Url    string `json:"url"`
	Name   string `json:"name"`
	Active *bool  `json:"active" gorm:"default:true"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
