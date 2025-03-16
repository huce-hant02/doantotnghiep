package model

import "time"

// Lớp học phần
type LoeCourseClass struct {
	ID       uint `json:"id"`
	CourseId uint `json:"courseId"`

	Code string `json:"code" gorm:"index:course_class_uniq_code,unique"`
	Year int    `json:"year"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	Students []LoeCourseClassStudent `json:"students" gorm:"foreignKey:CourseClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	OtherFilter LoeCourseClassOtherFilter `json:"otherFilter" gorm:"-"`

	// FK
	Course *Course `json:"course" gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LoeCourseClassOtherFilter struct {
	ListCourseId []int `json:"listCourseId"` // Course-IDs
}
