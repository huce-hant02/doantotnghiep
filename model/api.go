package model

import (
	"time"

	"gorm.io/gorm"
)

type API struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	Url         string `json:"url"`
	Method      string `json:"method"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      *bool  `json:"active" gorm:"default:true"`
	ForSync     *bool  `json:"forSync" gorm:"default:false"`
	AutoSync    *bool  `json:"autoSync" gorm:"default:false"`

	IsRunning *bool `json:"isRunning" gorm:"default:false"`

	DefaultParams  *string `json:"defaultParams"`
	DefaultPayload *string `json:"defaultPayload"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}
