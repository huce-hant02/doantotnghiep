package model

import (
	"time"
)

type Rubric struct {
	ID        uint   `json:"id"`
	OutlineID uint   `json:"outlineId"` //  gorm:"uniqueIndex:rubric_uniq"
	UUID      string `json:"uuid"`

	Code             string    `json:"code"` //  gorm:"uniqueIndex:rubric_uniq"
	Title            string    `json:"title"`
	IsClinicalCourse *bool     `json:"isClinicalCourse"`
	CreatedAt        time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt        time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt        gorm.DeletedAt `json:"-" swaggerignore:"true"`

	//foreignKey
	RubricItems []RubricItem `json:"rubricItems" gorm:"foreignKey:RubricID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
