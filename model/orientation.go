package model

import (
	"time"

	"gorm.io/gorm"
)

type Orientation struct {
	ID                 uint `json:"id"`
	EducationProgramID uint `json:"educationProgramID"`

	Code              string  `json:"code"`
	Title             string  `json:"title"`
	Index             *int    `json:"index"`
	ProgressTreeImage *string `json:"progressTreeImage"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

type OrientationRepository interface {
	Create(db *gorm.DB, orientations []Orientation, listOmitField []string) ([]Orientation, error)
	Update(db *gorm.DB, ID uint, orientation *Orientation) (*Orientation, error)
	Delete(db *gorm.DB, ID uint) error
	ISearchRepo
}
