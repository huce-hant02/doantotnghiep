package model

import (
	"time"
)

type CourseOrientation struct {
	ID              uint   `json:"id"`
	CourseID        uint   `json:"courseId"`
	OrientationID   uint   `json:"orientationId"`
	OrientationType string `json:"orientationType"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}
