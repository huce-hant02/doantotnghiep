package model

import (
	"time"

	"gorm.io/gorm"
)

type Faculty struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"userId" gorm:"default:0"`
	PuCode string  `json:"puCode" gorm:"uniqueIndex:faculty_pu_code_unique"`
	Email  *string `json:"email" gorm:"unique"`

	Code                string     `json:"code" gorm:"unique"`
	Name                string     `json:"name"`
	NameEng             string     `json:"nameEng"`
	AbbreviationNameEng string     `json:"abbreviationNameEng"`
	Type                *string    `json:"type"`
	SyncedAt            *time.Time `json:"syncedAt"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	Subjects []Subject `json:"subjects" gorm:"foreignKey:FacultyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Majors   []Major   `json:"majors" gorm:"foreignKey:DepartmentCode;references:PuCode;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`

	User *User `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
