package model

import (
	"time"

	"gorm.io/gorm"
)

type OutlineUpdateProcess struct {
	ID             uint   `json:"id"`
	OutlineID      uint   `json:"outlineId"`
	EmployeeID     uint   `json:"employeeId"`
	Description    string `json:"description"`
	NumberOfUpdate int    `json:"numberOfUpdate"`

	Employee Employee `swaggerignore:"true" json:"employee" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Outline  Outline  `swaggerignore:"true" json:"-" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
